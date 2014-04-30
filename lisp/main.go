package lisp

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	InitEnv()

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sexp := ReadSexp(in)
		fmt.Println("    read:", sexp)
		eval := GlobalEnv.Eval(sexp)
		fmt.Println("    eval:", eval)
	}
}
