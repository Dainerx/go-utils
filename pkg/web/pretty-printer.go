package web

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int

func PrintFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return err
	}

	VisitHtmlDoc(doc, startElement, endElement)
	return nil
}

func PrintFromUrl() {

}

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		if node.Data == "img" || node.Data == "a" {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", node.Data, fetchAttributes(node))
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		}
		depth++
	} else if node.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", node.Data)
		depth++
	} else if node.Type == html.CommentNode {
		fmt.Printf("%*s%s\n", depth*2, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	} else if node.Type == html.TextNode {
		depth--
	} else if node.Type == html.CommentNode {
		depth--
	}
}

func fetchAttributes(node *html.Node) string {
	var attributes string
	for _, a := range node.Attr {
		attributes += " " + a.Key + "=" + a.Val
	}
	return attributes
}
