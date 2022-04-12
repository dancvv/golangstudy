package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat().....")
}

func (h *Human) Walk() {
	fmt.Println("Human.walk()...")
}

type SuperMan struct {
	Human
	level int
}

func (s *SuperMan) Eat() {
	fmt.Println("SuoerMan.Eat()...")
}

// 子类的新方法
func (s *SuperMan) Fly() {
	fmt.Println("SuperMan flying")
}

func extendmain() {
	var s SuperMan
	s.name = "l21"
	s.level = 12
	s.sex = "male"

	s.Walk()
	s.Eat()
	s.Fly()
}
