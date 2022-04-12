package main

import "fmt"

// 定义一个函数，其返回类型也为string
type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println("something worked")
	fmt.Println(g(n))
}

func english(name string) string {
	return "hello," + name
}

func main() {
	// 通过如下方式将该函数转换成Greeting类型的函数对象
	// 传入english函数，将其转换为greeting
	greet := Greeting(english)
	greet.say("world")
}
