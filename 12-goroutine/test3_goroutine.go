package main

import "fmt"

func jkmain() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行。。。")
		c <- 666 // 将666发送给c
		fmt.Println("测试是否进行")
	}()
	num := <-c // 从c中接受数据，并赋值

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束。。。")

}
