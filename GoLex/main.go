package main

import (
	"nfa"
)

func main() {
	cmd := nfa.NewCommandLine()
	cmd.DoFile()
}
