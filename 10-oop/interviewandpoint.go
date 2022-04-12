package main

import "fmt"

type student struct {
	name string
	age  int
}

func testMain() {
	m := make(map[string]*student)
	stus := []student{
		{name: "aaa", age: 19},
		{name: "bbb", age: 20},
		{name: "ccc", age: 21},
	}
	for i, stu := range stus {
		// temp :=
		m[stu.name] = &stus[i]
		// fmt.Println(m)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.age)
	}
}
