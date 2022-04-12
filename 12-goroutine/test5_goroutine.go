package main

import "fmt"

func emain() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			// close(c)
		}
		// close可以关闭一个channel
		close(c)
	}()

	/*
		for {
			// ok如果为true表示channel没有关闭，如果为false表示channel关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	// 打印优化
	// 可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main finished...")
}
