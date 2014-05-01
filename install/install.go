package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"code.google.com/p/go.net/html"

	"github.com/golang/glog"
	"github.com/siteshen/go-practice/htmlhelper"
)

func checkError(err error) {
	if err != nil {
		glog.Fatalln(err)
	}
}

func MustGet(url string) io.ReadCloser {
	resp, err := http.Get(url)
	checkError(err)
	return resp.Body
}

func MustParse(url string) *html.Node {
	body := MustGet(url)
	node, err := html.Parse(body)
	checkError(err)

	filename := strings.Split(url, "/")
	f, err := os.Create(filename[len(filename)-1] + ".html")
	checkError(err)
	defer f.Close()

	html.Render(f, node)
	return node
}

// return a list for Repositories url
//
// http://godoc.org/code.google.com/p/go.tools
// http://godoc.org/code.google.com/p/go.crypto
// http://godoc.org/code.google.com/p/go.image
// ...
func SubRepoUrls(url string) (urls []string) {
	node := MustParse(url)
	h2list := htmlhelper.GetElementsByTagName(node, "h2")
	if h2list == nil {
		return
	}

	var repo *html.Node
	for _, node := range h2list {
		if strings.TrimSpace(htmlhelper.Text(node)) == "Repositories" {
			repo = NextSibling(node)
		}
	}

	hrefs := htmlhelper.GetElementsByTagName(repo, "a")
	for _, node := range hrefs {
		urls = append(urls, "http://godoc.org/"+htmlhelper.Text(node))
	}

	return
}

// SHIT: I don't want to get TextNode
func NextSibling(node *html.Node) (sibling *html.Node) {
	sibling = node.NextSibling
	for sibling != nil && sibling.Type == html.TextNode {
		sibling = sibling.NextSibling
	}
	return
}

func DirectoryURLs(url string) (urls []string) {
	node := MustParse(url)
	h3 := htmlhelper.GetElementById(node, "pkg-subdirectories")
	if h3 == nil {
		return
	}

	hrefs := htmlhelper.GetElementsByTagName(NextSibling(h3), "a")
	for _, node := range hrefs {
		urls = append(urls, "http://godoc.org"+htmlhelper.NodeAttr(node, "href"))
	}

	return
}

func ListRepos(baseUrl string) (urls []string) {
	urls = SubRepoUrls(baseUrl)
	for _, repo := range urls {
		fmt.Println(repo)
	}
	return
}

func ListDirectories(baseUrl string) (urls []string) {
	urls = DirectoryURLs(baseUrl)
	for _, u := range urls {
		fmt.Printf("\t%s\n", u)
	}
	return
}
