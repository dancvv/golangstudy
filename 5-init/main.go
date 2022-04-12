package main

import (
	"fmt"
	"time"
)

func mytest() {
	fmt.Println("hello , go")
}

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		// 避免因为执行过快，无法观察多线程执行效果
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()
	// go mytest()
	// time.Sleep(time.Second)
	// fmt.Println("hello, world")
	// // 阻塞main线程,使得其他线程执行完全
	// time.Sleep(time.Second)
	go mygo("协程1号")
	go mygo("协程2号")
	time.Sleep(time.Second)
}
