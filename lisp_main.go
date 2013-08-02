package main

import (
	"./lisp"
	"bufio"
	"fmt"
	"os"
)

func main() {
	env := lisp.InitEnv()
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sexp := lisp.ReadSexp(in)
		fmt.Println("    read:", sexp)
		eval := env.Eval(sexp)
		fmt.Println("    eval:", eval)
	}
}
