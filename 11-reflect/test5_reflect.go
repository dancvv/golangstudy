package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", this)
}

func lmain() {
	user := User{1, "sdf", 18}
	user.Call()
	DoFiledAndMethod(user)

}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType)

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is: ", inputValue)

	// 通过type获取里面的字段
	// 1.获取interface的type，通过type得到numfield，进行遍历
	// 2.得到每个field，数据类型
	// 3.通过filed有一个interface方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", filed.Name, filed.Type, value)
	}
	fmt.Println("------")
	// 通过type获取里面的方法，调用
	fmt.Println(inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
