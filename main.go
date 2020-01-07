package main

import (
	"fmt"
)

func main() {
	var f float32 = 16777216
	d := f + 2
	fmt.Printf("%v\n", d == f)
	fmt.Printf("hello, world\n")
}
