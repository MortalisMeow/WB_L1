package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Add(name string, age int) {
	h.name = name
	h.age = age
}

type Action struct {
	Human
}

func main() {

	act := Action{
		Human: Human{
			name: "Chel",
			age:  18,
		},
	}
	act.Add("Chelovek", 20)

	fmt.Println(act.name, act.age)
}
