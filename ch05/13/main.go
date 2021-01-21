package main

import (
	"fmt"
	"github/jongsuknim/gopl-exercises/ch05/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				err := download(item)
				if err != nil {
					fmt.Println(err)
				}
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func download(url string) error {
	domain, path, fname := parseUrl(url)

	fmt.Println(domain, path, fname)
	if domain == mainDomain {
		return nil
	}

	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	if path != "" {

		if err := createDir(dir + path); err != nil {
			return err
		}
	}

	return copyFile(url, dir+path+"/"+fname)
}

func createDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func copyFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Create %s", path)
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("Copy %s", path)
	}
	return nil
}

func parseUrl(url string) (domain, path, fname string) {
	path = ""
	fname = ""
	i := strings.Index(url, "://")
	if i > 0 {
		url = url[i+3:]
	}

	i1 := strings.Index(url, "/")
	if i1 < 0 {
		domain = url
		return
	}
	domain = url[:i1]

	i2 := strings.LastIndex(url, "/")
	if i2 < 0 {
		return
	}
	fname = url[i2+1:]

	if i1 == i2 {
		return
	}
	path = url[i1:i2]
	return
}

var mainDomain string

func main() {
	url := "https://www.google.com"
	domain, _, _ := parseUrl(url)
	mainDomain = domain

	breadthFirst(crawl, []string{url})
}
