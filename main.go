package main

import (
	"fmt"

	"github.com/Dainerx/go-utils/pkg/general"
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
	// err := web.PrintFromFile("pkg/web/short.html")
	// if err != nil {
	// 	fmt.Printf("PrintFromFile failed caused by: %v", err)
	// 	os.Exit(1)
	// }

	//fmt.Println(math.Max(2, 5, 8, 61, -1, 9))
	// fmt.Println(math.Min(2, 5, 8, 61, -1, 9))

	// s := strings.JoinVariant(";", "a", "ee", "bb")
	// fmt.Println(s)

	fmt.Println(general.ChangeReturnSquare(2, true))
}
