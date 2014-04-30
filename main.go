package main

import (
	"flag"

	"./htmlhelper"
	"./lisp"
	"./tour"
)

func main() {
	app := flag.String("name", "", "app you want to run: lisp | htmlhelpr | tour")
	flag.Parse()

	switch *app {
	case "lisp":
		lisp.Main()
	case "htmlhelper":
		htmlhelper.Main()
	case "tour":
		tour.Main()
	}
}
