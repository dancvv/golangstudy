package main

import "fmt"

// func swap(pa int, pb int) {
// 	var temp int
// 	temp = pa
// 	pa = pb
// 	pb = temp
// }

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
	fmt.Println("the value of ", temp)
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println(a, b)
	var p *int
	p = &a
	fmt.Println("the a's value")
	fmt.Println(&a)
	fmt.Println(a)
	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(pp)
}
