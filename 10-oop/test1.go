package main

import "fmt"

// 将一个类型定义为myint的别名
type myint int

// func main() {
// 	var a myint = 10
// 	fmt.Println("a = ", a)
// 	fmt.Println("type of a = %T", a)

// 	fmt.Println("------------")
// 	v1 := valueIntTest(1)
// 	fmt.Println(v1)
// 	v2 := pointerIntTest(&v1)
// 	fmt.Println(v2)
// 	fmt.Println("--------")
// 	structTestValue()
// }

func valueIntTest(a int) int {
	return a + 10
}

func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
}
