package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写，则该case就会进来
			t := x
			x = y
			y = t + y
		case data := <-quit:
			fmt.Println("quit", data)
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	// main go
	fibonacii(c, quit)
}
