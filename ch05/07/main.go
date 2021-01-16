package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func hasElementNodeChild(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			return true
		}
	}
	return false
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if hasElementNodeChild(n) {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		} else {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", strings.TrimSpace(n.Data))
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}

	for _, a := range n.Attr {
		fmt.Printf("%*s%s:%q", depth*2, "", a.Key, a.Val)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if hasElementNodeChild(n) {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("can't get from %s", url)
			continue
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Printf("parse error: %v", err)
			continue
		}

		forEachNode(doc, startElement, endElement)
	}
}
