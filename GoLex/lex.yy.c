#define YYPRIVATE

#line 1 "input.lex"
   int FCON = 1;
   int ICON = 2;


ifdef __NEVER__
/*------------------------------------------------
DFA (start state is 0) is :
 *
* State 0 [nonaccepting]

 * goto 2 on .
 * goto 4 on 0123456789
* State 1 [nonaccepting]

* State 2 [nonaccepting]

 * goto 1 on 0123456789
* State 3 [accepting, line 6, <  {printf("%s is a float number", yytext); return FCON;}>]

 * goto 3 on 0123456789
* State 4 [accepting, line 6, <  {printf("%s is a float number", yytext); return FCON;}>]

 * goto 3 on .
 * goto 5 on 0123456789
* State 5 [nonaccepting]

 * goto 2 on .
 * goto 5 on 0123456789
*/ 

#endif
/*
YY_TYPE 是宏定义，用于 DFA 状态转换表Yy_nxt[],它将会在下面进行定义。 宏定义YYF表示错误的状态跳转，当状态机跳转到错误状态时
模板代码会自动进行相应处理.DFA 状态机的起始状态为 0，同时宏定义 YYPRIVATE 也会在本模板文件中定义
*/
#ifndef YYPRIVATE
#define YYPRIVATE  static
#endif
#ifndef NULL
#include <stdio.h>
#endif
#include<debug.h>
#include<l.h>
#ifndef YYDEBUG
  int yydebug = 0;
#define YY_D(x) if(yydebug){x;}else
#else
#define YY_D(x)
#endif
typedef unsigned char YY_TYPE;
#define YYF  ((YYTYPE)(-1))
unsigned char* ii_text();

/*--------------------------------------
* The Yy_cmap[] and Yy_rmap arrays are used as follows:
* 
*  next_state= Yydtran[ Yy_rmap[current_state] ][ Yy_cmap[input_char] ];
* 
* Character positions in the Yy_cmap array are:
* 
*    ^@  ^A  ^B  ^C  ^D  ^E  ^F  ^G  ^H  ^I  ^J  ^K  ^L  ^M  ^N  ^O
*    ^P  ^Q  ^R  ^S  ^T  ^U  ^V  ^W  ^X  ^Y  ^Z  ^[  ^\  ^]  ^^  ^_
*         !   "   #   $   %   &   '   (   )   *   +   ,   -   .   /
*     0   1   2   3   4   5   6   7   8   9   :   ;   <   =   >   ?
*     @   A   B   C   D   E   F   G   H   I   J   K   L   M   N   O
*     P   Q   R   S   T   U   V   W   X   Y   Z   [   \   ]   ^   _
*     `   a   b   c   d   e   f   g   h   i   j   k   l   m   n   o
*     p   q   r   s   t   u   v   w   x   y   z   {   |   }   ~   DEL
*/

static unsigned char Yy_cmap[128]=
{
  0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,1,0,2,2,
    2,2,2,2,2,2,2,2,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0,0,0,
    0,0,0,0,0,0,0,0
};

static  unsigned char  Yy_rmap[6]=
{
    0,1,2,3,4,5
};

static unsigned char Yy_nxt[6][3]=
{
/*  0 */ {-1, 2, 4},
/*  1 */ {-1, -1, -1},
/*  2 */ {-1, -1, 1},
/*  3 */ {-1, -1, 3},
/*  4 */ {-1, 3, 5},
/*  5 */ {-1, 2, 5} 
};

/*--------------------------------------
* yy_next(state,c) is given the current state number and input
* character and evaluates to the next state.
*/

#define yy_next(state, c) (Yy_nxt[Yy_rmap[state]][Yy_cmap[c]])

/*--------------------------------------
 * 输出基于 DFA 的跳转表,首先我们将生成一个 Yyaccept数组，如果 Yyaccept[i]取值为 0，
	那表示节点 i 不是接收态，如果它的值不是 0，那么节点是接受态，此时他的值对应以下几种情况：
	1 表示节点对应的正则表达式需要开头匹配，也就是正则表达式以符号^开始，2 表示正则表达式需要
	末尾匹配，也就是表达式以符号$结尾，3 表示同时开头和结尾匹配，4 表示不需要开头或结尾匹配
 */

YYPRIVATE YY_TTYPE Yyaccept[]=
{
	0  ,  /*State 0  */
	0  ,  /*State 1  */
	0  ,  /*State 2  */
	4  ,  /*State 3  */
	4  ,  /*State 4  */
	0     /*State 5  */
};

/*--------------------------------------
*  语法解析器使用的全局变量放到这里，在词法解析器阶段，这里可以忽略
*/
char*  yytext;  /*指向当前正在读取的字符串*/
int    yyleng;  /*当前读取字符串的长度*/
int    yylineno;  /*当前读取字符串所在的行*/
FILE*  yyout = stdout;  /*默认情况下输入从控制台读取*/
#define output(c)  putc(c, yyout)
#define ECHO       fprintf(yyout, "%s", yytext)
#ifndef YYERROR
#define YYERROR  printf
#endif
#define  yymore() yymoreflg = 1
#define unput(c)   (ii_unput(c), --yyleng)
#define yyless(n) (ii_unterm(), yyleng -= ii_pushback(n) ? n : yyleng), ii_term() )
int input(void) {
    int c;
    if ((c = ii_input()) && (c != -1)) {
        yytext = (char*) ii_text();
        yylineno = ii_lineno();
        ++yyleng;
    }
    return c;
}
/*--------------------------------*/
yy_init_lex() {
    //做一些初始化工作，默认什么都不做
}
int yywrap() {
   //默认不要打开新文件
   return 0;
}
yylex() {
    int  yymoreflg;
    static  int yystate = -1;
    int  yylastaccept;
    int  yyprev;
    int yynstate;
    int yylook;
    int  yyanchor;
    if (yystate == -1) {
        //函数第一次执行时进入这里进行初始化
        yy_init_lex();
        ii_advance();
        ii_pushback(1);
    }
    yystate = 0;
    yylastaccept = 0;
    yymoreflg = 0;
    ii_unterm();
    ii_mark_start();
    while(1) {
        /*
        首先检测当前读入的文件是否已经到了末尾。如果是，并且当前有没有处理的接收状态,yylastaccept 的值就不是 0
        ，那么此时就先执行对应接收状态的代码，如果在到了文件末尾还没有遇到过接收状态，那么尝试打开新的输入文件，
        如果新文件打开失败则返回
        */
        while(1) {
            if((yylook = ii_look(1) != EOF) {
                yynstate = yy_next(yystate, yylook);
                break;
            } else {
                if (yylastaccept) {
                    yynstate = YYF;
                    break;
                } else if (yywrap()) {
                    //yywrap 打开新的输入文件,进入到这里说明没有新的文件要打开
                    yytext = "";
                    yyleng = 0;
                    return 0;
                } else {
                    //这里说明打开了新的输入文件
                    ii_advance(); //读取数据到缓冲区
                    ii_pushback(1);
                }
            }
        }
        if (yynstate != YYF) {
            YY_D(printf("    Transition from state %d", yystate));
            YY_D(printf(" to state %d on <%c>\n", yystate, yylook));
            if (ii_advance() < 0) {
                YYERROR("Line %d, lexeme too long. Discarding extra characters.\n", ii_lineno());
                ii_flush(1);
            }
            if (yyanchor = Yyaccept[yynstate]) {
                //当前状态是接收状态
                yyprev = yystate;
                yylastaccept = yynstate;
                ii_mark_end();
            }
            yystate = yynstate;
        } else {
            //在这里意味着当前状态机接收字符后进入错误状态，于是我们处理之前进入的接收状态
            if (!yylastaccept) {
                //此前没有进入过接收状态
                YYERROR("Ignoring bad input\n");
                ii_advance();
            } else {
                //处理之前进入的接收状态
                ii_to_mark();
                if (yyancor & 2) {
                    //末尾匹配，先将当前回车字符放回缓冲区
                    ii_pushback(1);
                }
                if (yyanchor & 1) {
                    //开头匹配，忽略掉当前输入字符串开头的回车字符
                    ii_move_start();
                }
                ii_term();  //将当前输入字符串的末尾添加\0 符号
                yytext = (char*)ii_text();
                yyleng = ii_length();
                yylineno = ii_lineno();
                YY_D(printf("Accepting state %d, ", yylastaccept);
                YY_D(printf("line %d: <%s>\n", yylineno, yytext));
                switch(yylastaccept) {
		case 3:					/* State 3   */
		      {printf("%s is a float number", yytext); return FCON;}
		    break;
		case 4:					/* State 4   */
		      {printf("%s is a float number", yytext); return FCON;}
		    break;
                default:
                    YYERROR("INTERNAL ERROR, yylex: Unkonw accept state %d.\n", yylastaccept);
                    break;
                }
            }
            ii_unterm();
            yylastaccept = 0;
            if (!yymoreflg) {
                yystate = 0;
                ii_mark_start();
            } else {
                yystate = yyprev; //记录上一次遇到的状态
                yymoreflg = 0;
            }
        }
    }
}
