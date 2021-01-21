package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func join(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func forEachNode(n *html.Node, f func(n *html.Node) bool) []*html.Node {
	if f != nil && f(n) {
		return []*html.Node{n}
	}
	retval := make([]*html.Node, 0)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r := forEachNode(c, f)
		retval = append(retval, r...)
	}
	return retval
}

func exist(lis []string, d string) bool {
	for _, v := range lis {
		if v == d {
			return true
		}
	}
	return false
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return forEachNode(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && exist(name, n.Data)
	})
}

func main() {
	url := "https://www.gopl.io/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("getting %s: %s", url, resp.Status)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("parsing %s as HTML: %v", url, err)
		os.Exit(1)
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Println(images)
	fmt.Println(headings)
}
