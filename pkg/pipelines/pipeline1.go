package pipelines

// func main() {

// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			naturals <- i
// 		}
// 		close(naturals)
// 	}()

// 	go func() {
// 		for {
// 			x, ok := <-naturals
// 			if !ok {
// 				break
// 			}
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()

// 	for {
// 		x, ok := <-squares
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(x)
// 	}

// }
