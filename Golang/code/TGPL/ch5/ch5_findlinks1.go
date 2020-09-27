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

/*
5.2 编写函数，记录在HTML树中出现的同名元素的次数。
*/

func countElementTag(cnt map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		eleTag := n.Data
		cnt[eleTag] += 1
	}

	// change call `visit` from loop to recursion
	recursionNode := n.FirstChild
	if recursionNode != nil {
		cnt = countElementTag(cnt, recursionNode)
	}
	recursionNode = n.NextSibling
	if recursionNode != nil {
		cnt = countElementTag(cnt, recursionNode)
	}
	//fmt.Println(cnt)
	return cnt
}

func testCountElementTag() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	cnt := make(map[string]int)
	result := countElementTag(cnt, doc)
	for k, v := range result {
		fmt.Println(k, v)
	}
}
