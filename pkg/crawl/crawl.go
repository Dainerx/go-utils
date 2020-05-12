// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item to return the links
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // set of strings
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...) //append all links returned by f = crawl
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

func discover1() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

// Unbounded parallel
func discover2() {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for w := range worklist {
		for _, item := range w {
			if _, ok := seen[item]; !ok {
				seen[item] = true
				go func(url string) {
					worklist <- crawl(url)
				}(item)
			}
		}
	}

}

func discover3() {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for w := range worklist {
		for _, item := range w {
			if _, ok := seen[item]; !ok {
				seen[item] = true
				go func(url string) {
					worklist <- crawl(url)
				}(item)
			}
		}
	}

}

//!+main
func main() {
	discover2()
}
