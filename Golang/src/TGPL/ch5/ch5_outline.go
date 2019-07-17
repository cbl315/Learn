package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

// outline 无出栈操作 当stack的元素数量过大导致底层数组改变时， 调用方的stack不会被修改
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}

func testOutline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}
	outline(nil, doc)
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrStr := ""
		defer func() { depth++ }()
		for _, attr := range n.Attr {
			k, v := attr.Key, attr.Val
			attrStr += fmt.Sprintf(" %s='%s'", k, v)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrStr)
				return
			}
		}
		fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrStr)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
				return
			}
		}
	}
}

func testOutline2() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}
	forEachNode(doc, startElement, endElement)
}
