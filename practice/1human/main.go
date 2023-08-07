package main

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).*/
import "fmt"

// Human struct parent
type Human struct {
	age    int
	gender string
}

// Parent methods
func (h *Human) getAge() int {
	return h.age
}

func (h *Human) getGender() string {
	return h.gender
}

// Child struct
type Action struct {
	Human
}

func main() {
	action := Action{Human{
		age:    20,
		gender: "men",
	}}
	fmt.Println(action.getAge())
	fmt.Println(action.getGender())
}
