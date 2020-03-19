package main

import (
	"fmt"
	"os"

	"github.com/Dainerx/go-utils/pkg/web"
)

func main() {

	// if err := web.FindLinksInHtmlFile("pkg/web/golang.org.html"); err != nil {
	// 	fmt.Printf("FindLinksInHtmlFile failed caused by: %v", err)
	// 	os.Exit(1)
	// }

	// counter, err := web.CountElementsInHtmlFile("pkg/web/golang.org.html")
	// if err != nil {
	// 	fmt.Printf("FindLinksInHtmlFile failed caused by: %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(counter)

	// words, images, err := web.CountWordsAndImages("https://golang.org")
	// if err != nil {
	// 	fmt.Printf("CountWordsAndImages failed caused by: %v", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("words = %d, images = %d\n", words, images)
	err := web.PrintFromFile("pkg/web/short.html")
	if err != nil {
		fmt.Printf("PrintFromFile failed caused by: %v", err)
		os.Exit(1)
	}

}
