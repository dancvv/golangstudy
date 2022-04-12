package main

import "fmt"

func deferFun1() {
	fmt.Println("A")
}

func deferFun2() {
	fmt.Println("B")
}

func deferFun3() {
	fmt.Println("C")
}

/**
defer 与 return 方法的测试，谁先执行
*/

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func retrunFunc() int {
	fmt.Println("return func called...")
	return 0
}

func deferAreturn() int {
	defer deferFunc()
	return retrunFunc()
}

func main() {
	// defer deferFun1()
	// defer deferFun2()
	// defer deferFun3()

	// 测试defer和return的区别
	deferAreturn()

}
