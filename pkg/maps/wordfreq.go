package maps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordFreq(input string) map[string]int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	wordsFreq := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		wordsFreq[word]++
	}
	return wordsFreq
}

func WordFreqFromFile(filePath string) map[string]int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Input file could not be read")
		return nil
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordsFreq := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		wordsFreq[word]++
	}
	return wordsFreq
}
