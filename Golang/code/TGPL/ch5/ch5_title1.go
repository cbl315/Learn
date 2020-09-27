package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err!= nil {
		return err
	}
	// check Content-Type is whether HTML (e.g, "text/html;charset=utf-8")
	ct := resp.Header.Get("Content-Type")
	defer resp.Body.Close()
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func testTitle() {
	err := title("https://ss2.bdstatic.com/kfoZeXSm1A5BphGlnYG/skin/1.jpg?2&quot")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}


// soleTitle returns the text of the first non-empty title elemetn
// in doc, and an error if there was not exactly one
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct {}
	defer func() {
		switch p := recover(); p{
		case nil : // no panic
		case bailout{}: // expected panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
			}
	}()
	// Bail out of recursion if we find more than one nonempty title
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}