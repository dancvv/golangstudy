package main

import (
	"fmt"
	"reflect"
)

// 反射基本类型

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func kmain() {
	var num float64 = 1.2345
	reflectNum(num)
}
