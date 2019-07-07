// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// change call `visit` from loop to recursion
	recursionNode := n.FirstChild
	if recursionNode != nil {
		links = visit(links, recursionNode)
	}
	recursionNode = n.NextSibling
	if recursionNode != nil {
		links = visit(links, recursionNode)
	}

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
	return links
}

func testFindlinks1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
