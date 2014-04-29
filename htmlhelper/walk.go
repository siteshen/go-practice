package htmlhelper

import "code.google.com/p/go.net/html"

// helper for walker
func True(*html.Node) bool  { return true }
func False(*html.Node) bool { return false }

type Predictor func(*html.Node) bool
type Processor func(*html.Node)

func walk(node *html.Node, f Processor, pred Predictor) {
	f(node)
	if !pred(node) {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			walk(child, f, pred)
		}
	}
}

func WalkNodeUntil(node *html.Node, f Processor, pred Predictor) {
	walk(node, f, pred)
}

func WalkNode(node *html.Node, f Processor) {
	walk(node, f, False)
}

func FilterNode(node *html.Node, pred Predictor) (nodes []*html.Node) {
	var Collect = func(n *html.Node) {
		if pred(n) {
			nodes = append(nodes, n)
		}
	}

	walk(node, Collect, False)
	return
}
