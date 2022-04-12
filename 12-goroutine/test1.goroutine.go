package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}

func amain() {
	// 创建一个go程， 去执行newTask流程
	go newTask()
	fmt.Println("main goroutine exit")

	// i := 0
	// // for package main

	// for {
	// 	i++
	// 	fmt.Printf("main goroutine: i = %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
