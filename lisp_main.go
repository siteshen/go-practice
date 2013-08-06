package main

import (
	"./lisp"
	"bufio"
	"fmt"
	"os"
)

func main() {
	lisp.InitEnv()

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sexp := lisp.ReadSexp(in)
		fmt.Println("    read:", sexp)
		eval := lisp.GlobalEnv.Eval(sexp)
		fmt.Println("    eval:", eval)
	}
}
