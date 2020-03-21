package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Dainerx/go-utils/pkg/maps"

	"golang.org/x/net/html"
)

// Visit a html node recursively, calls pre(node) and post(node) for every node.
// pre - preorder is called before visiting a child's node.
// post - postorder is called after visiting a child's node.
func VisitHtmlDoc(node *html.Node, pre, post func(node *html.Node)) {
	if node != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		VisitHtmlDoc(c, pre, post)
	}

	if node != nil {
		post(node)
	}
}

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

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return //bare return: return 0,0,err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return //another bare return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(node *html.Node) (words, images int) {
	if node.Type == html.TextNode {
		// Get map of distinct words and their count of one node.
		for _, wordCount := range maps.WordFreq(node.Data) {
			words += wordCount
		}
	} else if node.Type == html.ElementNode && node.Data == "img" {
		images++
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		wordCount, imageCount := countWordsAndImages(c)
		words += wordCount
		images += imageCount
	}
	return words, images
}

func ElementByID(doc *html.Node, id string) *html.Node {
	// trivial do same as visit test for attribute key id.
}
