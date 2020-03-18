package main

import (
	"fmt"
	"os"

	"github.com/Dainerx/go-utils/pkg/web"
)

func main() {

	if err := web.FindLinksFromHtmlFile("pkg/web/golang.org.html"); err != nil {
		fmt.Printf("Failed caused by: %v", err)
		os.Exit(1)
	}
}
