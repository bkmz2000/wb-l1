package main

import "fmt"

type Human struct {
	age  int
	name string
}

func (h *Human) Init(age int, name string) {
	h.age = age
	h.name = name
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) celebrateBirthday() {
	fmt.Printf("Happy birthday, %s!", h.name)
	h.age += 1
}

type Action struct {
	Human
	Object     Human
	ActionType int
}
