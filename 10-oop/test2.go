package main

import "fmt"

type PersonD struct {
	id   int
	name string
}

// 接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

// 接收者为指针类型
func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func strucTestFunc() {
	// 与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用

	// 值类型调用方法
	personValue := PersonD{101, "hello world"}
	fmt.Print(personValue)
	personValue.pointShowName()
	personValue.valueShowName()

	// 指针类型调用方法
	personPointer := &PersonD{102, "hello golang"}
	personPointer.valueShowName()
	personPointer.pointShowName()
}

// func main() {
// 	strucTestFunc()
// }
