package htmlhelper

import "code.google.com/p/go.net/html"

// document.GetElementById()
func GetElementById(node *html.Node, id string) (result *html.Node) {
	pred := func(n *html.Node) bool { return NodeAttr(n, "id") == id }
	walker := func(n *html.Node) {
		if pred(n) {
			result = n
		}
	}

	WalkNodeUntil(node, walker, pred)
	return
}

// document.GetElementById()
func GetElementsByName(node *html.Node, name string) []*html.Node {
	pred := func(n *html.Node) bool {
		return n.Type == html.ElementNode && NodeAttr(n, "name") == name
	}

	return FilterNode(node, pred)
}

// node.GetElementsByTagName()
func GetElementsByTagName(node *html.Node, tag string) []*html.Node {
	pred := func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}

	return FilterNode(node, pred)
}

// node.GetElementsByClassName()
func GetElementsByClassName(node *html.Node, cls string) []*html.Node {
	pred := func(n *html.Node) bool {
		for _, class := range NodeClasses(n) {
			if class == cls {
				return true
			}
		}
		return false
	}

	return FilterNode(node, pred)
}
