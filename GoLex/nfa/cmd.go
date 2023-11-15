package nfa

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type CommandLine struct {
	lexerReader  *LexReader
	parser       *RegParser
	nfaConverter *NfaDfaConverter
	//打开 lex.par 文件对应的句柄
	tempFile *os.File
	//记录lex.par 读到第几行
	inputLines int
	scanner    *bufio.Scanner
}

func NewCommandLine() *CommandLine {
	lexReader, _ := NewLexReader("input.lex", "lex.yy.c")
	parser, _ := NewRegParser(lexReader)
	nfaConverter := NewNfaDfaConverter()
	nfaConverter.fp = lexReader.OFile

	return &CommandLine{
		lexerReader:  lexReader,
		parser:       parser,
		nfaConverter: nfaConverter,
	}
}

func (c *CommandLine) head() {
	//先读取 input.lex ，识别其中的正则表达式
	c.lexerReader.Head()
}

func (c *CommandLine) miniDFA() {
	//构建 DFA 状态机并最小化状态机
	start := c.parser.Parse()
	c.nfaConverter.MakeDTran(start)
	c.nfaConverter.PrintDfaTransition()
	c.nfaConverter.MinimizeDFA()

	fmt.Printf("%d out of DFA states in miminized machine\n", c.nfaConverter.nstates,
		DFA_MAX)

}

func (c *CommandLine) PrintHeader() {
	//针对未压缩的 DFA 状态就，输出对应的 c 语言注释
	c.PrintUnCompressedDFA()
}

func (c *CommandLine) Signon() {
	//这里设置当前时间
	date := time.Now()
	//这里设置你的名字
	name := "yichen"
	fmt.Printf("GoLex 1.0 [%s] . (c) %s, All rights reserved\n", date.Format("01-02-2006"), name)
}

func (c *CommandLine) PrintUnCompressedDFA() {
	fmt.Fprint(c.nfaConverter.fp, "#ifdef __NEVER__\n")
	fmt.Fprint(c.nfaConverter.fp, "/*------------------------------------------------\n")
	fmt.Fprint(c.nfaConverter.fp, "DFA (start state is 0) is :\n *\n")
	nrows := c.nfaConverter.nstates
	charsPrinted := 0
	for i := 0; i < nrows; i++ {
		dstate := c.nfaConverter.dstates[i]
		if dstate.isAccepted == false {
			fmt.Fprintf(c.nfaConverter.fp, "* State %d [nonaccepting]\n", dstate.state)
		} else {
			//这里需要输出行数
			//fmt.Fprintf(n.fp, "* State %d [accepting, line %d <", i, )
			fmt.Fprintf(c.nfaConverter.fp, "* State %d [accepting, line %d, <%s>]\n", i, dstate.LineNo, dstate.acceptString)
			if dstate.anchor != NONE {
				start := ""
				end := ""
				if (dstate.anchor & START) != NONE {
					start = "start"
				}
				if (dstate.anchor & END) != NONE {
					end = "end"
				}
				fmt.Fprintf(c.nfaConverter.fp, " Anchor: %s %s \n", start, end)
			}
		}
		lastTransition := -1
		for j := 0; j < MAX_CHARS; j++ {
			if c.nfaConverter.dtrans[i][j] != F {
				if c.nfaConverter.dtrans[i][j] != lastTransition {
					fmt.Fprintf(c.nfaConverter.fp, "\n * goto %d on ", c.nfaConverter.dtrans[i][j])
					charsPrinted = 0
				}
				fmt.Fprintf(c.nfaConverter.fp, "%s", c.nfaConverter.BinToAscii(j))
				charsPrinted += len(c.nfaConverter.BinToAscii(j))
				if charsPrinted > 56 {
					//16 个空格
					fmt.Fprintf(c.nfaConverter.fp, "\n *                ")
					charsPrinted = 0
				}
				lastTransition = c.nfaConverter.dtrans[i][j]
			}
		}
		fmt.Fprintf(c.nfaConverter.fp, "\n")
	}
	fmt.Fprintf(c.nfaConverter.fp, "*/ \n\n")
	fmt.Fprintf(c.nfaConverter.fp, "#endif\n")
}

func (c *CommandLine) PrintDriver() {
	text := "输出基于 DFA 的跳转表,首先我们将生成一个 Yyaccept数组，如果 Yyaccept[i]取值为 0，" +
		"\n\t那表示节点 i 不是接收态，如果它的值不是 0，那么节点是接受态，此时他的值对应以下几种情况：" +
		"\n\t1 表示节点对应的正则表达式需要开头匹配，也就是正则表达式以符号^开始，" +
		"2 表示正则表达式需要\n\t末尾匹配，也就是表达式以符号$结尾，3 表示同时开头和结尾匹配，4 表示不需要开头或结尾匹配"
	comments := make([]string, 0)
	comments = append(comments, text)
	c.nfaConverter.comment(comments)
	//YYPRIVATE YY_TTYPE 是 c 语言代码中的宏定义，我们将在后面代码提供其定义
	//YYPRIVATE 对应 static, YY_TTYPE 对应 unsigned char
	fmt.Fprintf(c.nfaConverter.fp, "YYPRIVATE YY_TTYPE Yyaccept[]=\n")
	fmt.Fprintf(c.nfaConverter.fp, "{\n")
	for i := 0; i < c.nfaConverter.nstates; i++ {
		if c.nfaConverter.dstates[i].isAccepted == false {
			//如果节点i 不是接收态，Yyaccept[i] = 0
			fmt.Fprintf(c.nfaConverter.fp, "\t0  ")
		} else {
			anchor := 4
			if c.nfaConverter.dstates[i].anchor != NONE {
				anchor = int(c.nfaConverter.dstates[i].anchor)
			}
			fmt.Fprintf(c.nfaConverter.fp, "\t%-3d", anchor)
		}

		if i == c.nfaConverter.nstates-1 {
			fmt.Fprint(c.nfaConverter.fp, "   ")
		} else {
			fmt.Fprint(c.nfaConverter.fp, ",  ")
		}
		fmt.Fprintf(c.nfaConverter.fp, "/*State %-3d*/\n", i)
	}
	fmt.Fprintf(c.nfaConverter.fp, "};\n\n")
	// 先处理 switch cases 之前的内容
	c.driver_2()
	//如果遇到接收状态，我们将接收状态对应的操作代码拷贝到switch(yylastaccept)里面的 case 分支
	for i := 0; i < c.nfaConverter.nstates; i++ {
		if c.nfaConverter.dstates[i].isAccepted {
			fmt.Fprintf(c.nfaConverter.fp, "\t\tcase %d:\t\t\t\t\t/* State %-3d */\n", i, i)
			fmt.Fprintf(c.nfaConverter.fp, "\t\t    %s\n", c.nfaConverter.dstates[i].acceptString)
			fmt.Fprint(c.nfaConverter.fp, "\t\t    break;\n")
		}
	}

	//输出 switch case 之后内容
	c.driver_2()
}

func (c *CommandLine) driver_1() {
	/*
				打开名为 lex.par 的代码模板文件，这里我们将模板文件名字写死为 lex.par 是为了简单，
			    它的名字完全可以设计成从命令行输入,driver_1 函数的作用主要是打开文件，
		        真正实现内容拷贝操作的是在 driver_2
	*/
	tempFile, err := os.Open("lex.par")
	if err != nil {
		panic("error for open lex.par file")
	}
	c.inputLines = 0
	c.tempFile = tempFile
	c.scanner = bufio.NewScanner(tempFile)
	//开始拷贝模板内容
	c.driver_2()

}

func (c *CommandLine) readTempFile() (string, bool) {
	res := c.scanner.Scan()
	textLine := ""
	if res == true {
		textLine = c.scanner.Text()
	}

	return textLine, res
}

func (c *CommandLine) driver_2() bool {
	processingComment := false
	textLine, endOfFile := c.readTempFile()
	for endOfFile != false {
		if len(textLine) == 0 {
			textLine, endOfFile = c.readTempFile()
			continue
		}
		c.inputLines += 1
		if textLine[0] == '\f' {
			/*
				读取到符号\f 表示 lex.par 中一个特定区域的内容已经读取完毕
			*/
			break
		}

		idx := 0
		for textLine[idx] == ' ' {
			//忽略掉空格
			idx += 1
		}
		if textLine[idx] == '@' {
			//读取到@表示当前模板内容属于注释
			processingComment = true
			textLine, endOfFile = c.readTempFile()
			continue
		} else if processingComment {
			//上一行是注释，当前行不是
			processingComment = false
		}
		//把当前读到的模板内容写入 lex.yy.c
		fmt.Fprintf(c.nfaConverter.fp, "%s\n", textLine)

		textLine, endOfFile = c.readTempFile()
	}

	return endOfFile
}

func (c *CommandLine) pnext(name string) {
	/*
		这里输出一段c语言代码，它们的作用是根据输入字符，查询上面生成的压缩跳转表后，跳转到下一个状态节点
	*/

	toptext := []string{"unsigned int c ;", "int cur_state;", "{",
		"/*给定当前状态和输入字符，返回跳转后的下一个状态节点*/",
	}
	boptext := []string{
		"    int i;",
		"",
		"    if ( p ) ",
		"    {",
		"        if ((i = *p++) == 0)",
		"            return p[ c ];",
		"",
		"        for(; --i >= 0; p += 2)",
		"            if( c == p[0]  )",
		"                return p[1];",
		"    }",
		"    return unsigned char (-1);",
		"}",
	}

	fmt.Fprintf(c.nfaConverter.fp, "\n/*------------------------------------*/\n")
	fmt.Fprintf(c.nfaConverter.fp, "%s %s yynext(cur_state, c)\n", SCLASS, TYPE)
	for _, str := range toptext {
		fmt.Fprintf(c.nfaConverter.fp, "%s\n", str)
	}
	fmt.Fprintf(c.nfaConverter.fp, "    %s    *p = %s[cur_state];\n", TYPE, name)
	for _, str := range boptext {
		fmt.Fprintf(c.nfaConverter.fp, "%s\n", str)
	}
}

func (c *CommandLine) cnext(name string) {
	text := []string{
		"yy_next(state,c) is given the current state number and input",
		"character and evaluates to the next state.",
	}
	c.nfaConverter.printComment(text)
	fmt.Fprintf(c.nfaConverter.fp, "#define yy_next(state, c) (%s[Yy_rmap[state]][Yy_cmap[c]])\n", name)
}

func (c *CommandLine) tail() {
	//忽略掉含有 %% 的行
	c.readTempFile()
	//将%% 后面的内容直接拷贝到 lex.yy.c
	textLine, endOfFile := c.readTempFile()
	for endOfFile != false {
		fmt.Printf("t%d: %s", c.inputLines, textLine)
		c.inputLines += 1
		fmt.Fprintf(c.nfaConverter.fp, "%s\n", textLine)
	}

	c.nfaConverter.fp.Close()
}

func (c *CommandLine) DoFile() {
	c.head()
	c.miniDFA()
	c.PrintHeader()
	c.driver_1()
	//压缩 DFA 状态机
	c.nfaConverter.DoSquash()
	c.cnext("Yy_nxt")
	//输出跳转表
	c.PrintDriver()
	c.tail()
}
