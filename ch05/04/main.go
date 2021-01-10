package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func getValue(attr []html.Attribute) string {
	for _, a := range attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		switch n.Data {
		case "a":
			if l := getValue(n.Attr); l != "" {
				links = append(links, l)
			}
		case "img":
			if l := getValue(n.Attr); l != "" {
				links = append(links, l)
			}
		case "script":
			if l := getValue(n.Attr); l != "" {
				links = append(links, l)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
