package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const regexEanLine = `"ean"`

func isEanLine(line string) bool {
	matched, err := regexp.MatchString(regexEanLine, line)
	if err != nil {
		return false
	}

	return matched
}

func getEanFromLine(eanLine string) string {
	var ean string
	c := 0
	numbers := false
	for _, r := range eanLine {
		if r == '"' {
			if numbers {
				numbers = false
			}
			c++
		}

		if numbers {
			ean += string(r)
		}

		if c == 3 {
			numbers = true
		}

	}

	return ean
}

func main() {
	source, err := os.Open("seed.json")
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, err := os.Create("seed-cleaned.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		line := scanner.Text()
		if isEanLine(line) {
			eanNumber := getEanFromLine(line)
			if len(eanNumber) == 0 {
				continue
			}

			fmt.Fprintln(destination, eanNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
