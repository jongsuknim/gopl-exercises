package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var stopped bool

func forEachNode2(n *html.Node, pre, post func(n *html.Node) bool) {
	if stopped {
		return
	}

	if pre != nil {
		if pre(n) {
			stopped = true
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	stopped = false
	forEachNode2(n, pre, post)
}

var foundNode *html.Node

func isId(id string) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				foundNode = n
				return true
			}
		}
		return false
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	forEachNode(doc, isId(id), isId(id))
	return foundNode
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

		ElementByID(doc, "suggestions")
		if foundNode != nil {
			fmt.Printf("found:%q", *foundNode)
		}
	}
}
