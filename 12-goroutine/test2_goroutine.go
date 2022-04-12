package main

import (
	"fmt"
	"time"
)

func lmain() {
	// 用go创建承载一个形参为空，返回值为空的一个函数
	// go func() {
	// 	defer fmt.Println("A.defer")

	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		runtime.Goexit()
	// 		// 退出当前goroutine
	// 		fmt.Println("B")
	// 	}()

	// 	fmt.Println("A")
	// }()

	// 调用形参的函数
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
