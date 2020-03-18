package web

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func FindLinksFromStandardInput() error {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		return err
	}
	visit(doc)
	return nil
}

func FindLinksInHtmlFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return err
	}

	visit(doc)
	return nil

}

func visit(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

func CountElementsInHtmlFile(filePath string) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	elementCounter := make(map[string]int)
	visitCount(elementCounter, doc) // map is passed by reference.
	return elementCounter, nil
}

func visitCount(elementCounter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elementCounter[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitCount(elementCounter, c)
	}
}
