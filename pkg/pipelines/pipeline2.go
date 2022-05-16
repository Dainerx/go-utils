package pipelines

import "fmt"

// unidirectional channel for sending only
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}

	close(out)
}

// send only / read only
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}

	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)
// 	go counter(naturals)
// 	go squarer(squares, naturals)
// 	printer(squares)
// }
