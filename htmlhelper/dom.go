package htmlhelper

import (
	"bytes"
	"strings"

	"code.google.com/p/go.net/html"
)

func NodeAttr(node *html.Node, name string) string {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == name {
				return attr.Val
			}
		}
	}
	return ""
}

func NodeClasses(node *html.Node) (classes []string) {
	return strings.Split(NodeAttr(node, "class"), " ")
}

func NodeText(node *html.Node) (text string) {
	if node.Type == html.TextNode {
		text = node.Data
	}
	return
}

func Html(node *html.Node) string {
	var buffer bytes.Buffer
	html.Render(&buffer, node)
	return buffer.String()
}

func Text(node *html.Node) string {
	texts := []string{}
	walker := func(n *html.Node) {
		texts = append(texts, NodeText(n))
	}
	WalkNode(node, walker)
	return strings.Join(texts, "")
}
