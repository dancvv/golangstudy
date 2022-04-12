package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

// 子类实现了父类的所有接口

func (c *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "CAT"
}

// 父类类型的变量指向子类的具体数据变量

func main() {
	var animal AnimalIF
	animal = &Cat{"green"}
	animal.Sleep()
}
