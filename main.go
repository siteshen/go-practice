package main

import (
	"flag"
	"fmt"
	"strings"

	"./bookmark"
	"./htmlhelper"
	"./lisp"
	"./tour"
)

var AppList = map[string]func(){
	"lisp":       lisp.Main,
	"tour":       tour.Main,
	"bookmark":   bookmark.Main,
	"htmlhelper": htmlhelper.Main,
}

func appNames() (names []string) {
	for name, _ := range AppList {
		names = append(names, name)
	}
	return
}

func main() {
	var appname string
	usage := fmt.Sprintf("app you want to run: %s", strings.Join(appNames(), " | "))
	flag.StringVar(&appname, "appname", "", usage)
	flag.Parse()

	f := AppList[appname]
	if f == nil {
		flag.Usage()
		return
	}
	f()
}
