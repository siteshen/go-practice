package htmlhelper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"code.google.com/p/go.net/html"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Test(r io.Reader) {
	node, err := html.Parse(r)
	checkError(err)

	footer := GetElementById(node, "x-footer")
	fmt.Printf("GetElementById: \n%+v\n\n", Text(footer))
	// html.Render(os.Stdout, footer)

	metas := GetElementsByName(node, "viewport")
	for _, elem := range metas {
		fmt.Printf("GetElementsByName: \n%+v\n", Html(elem))
	}

	elems := GetElementsByClassName(node, "clearfix")
	for _, elem := range elems {
		fmt.Printf("GetElementsByClassName: \n%+v\n", Text(elem))
	}

	elems = GetElementsByTagName(node, "button")
	for _, elem := range elems {
		fmt.Printf("GetElementsByTagName: \n%+v\n", Html(elem))
	}
}

func Wget(url, filename string) (err error) {
	log.Printf("%s fetching...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	log.Printf("%s saving...\n", url)
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	log.Printf("%s saved.", url)

	return
}

func ParseHtml(in io.Reader) {
	node, err := html.Parse(in)
	checkError(err)

	fmt.Println(Html(GetElementById(node, "x-footer")))
}

func Main() {
	// fetch content
	// Wget("http://godoc.org/-/subrepo", "sub.html")

	// parse content
	file, err := os.Open("sub.html")
	checkError(err)
	defer file.Close()

	// Test(file)
	ParseHtml(file)
}
