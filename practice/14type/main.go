package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	var val interface{}
	val = make(chan int)
	fmt.Println(fmt.Sprintf("%T", val)) // "[]int"
}
