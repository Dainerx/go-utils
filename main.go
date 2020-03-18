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

	counter, err := web.CountElementsInHtmlFile("pkg/web/golang.org.html")
	if err != nil {
		fmt.Printf("FindLinksInHtmlFile failed caused by: %v", err)
		os.Exit(1)
	}
	fmt.Println(counter)
}
