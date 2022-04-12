package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	level int
}

func (h *Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.level)
	fmt.Println("-------")
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func (h *Hero) SetAge(newAd int) {
	h.Ad = newAd
}

func sonmain() {
	hero := Hero{Name: "zhang3", Ad: 100}
	hero.Show()

	hero2 := Hero{level: 12}
	fmt.Println(hero2.level)
	hero2.SetName("li4")
	hero2.SetAge(15)
	hero2.Show()
}
