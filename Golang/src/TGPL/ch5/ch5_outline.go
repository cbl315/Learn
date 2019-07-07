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
