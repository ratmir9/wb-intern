package main

import "fmt"

// Новый тип Human
type Human struct {
	Name    string
	Surname string
}

// функция приветствия
func (h *Human) SayHello() string {
	return fmt.Sprintf("Меня зовут %s %s", h.Surname, h.Name)
}

type Action struct {
	Human // встроенный тип Human
}

func main() {
	a := &Action{
		Human: Human{
			Name:    "Ivan",
			Surname: "Ivanov",
		},
	}
	fmt.Println(a.SayHello())
}
