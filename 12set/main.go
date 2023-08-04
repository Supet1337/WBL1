package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	maptemp := make(map[string]struct{})
	set := make([]string, 0)

	for _, el := range str {
		maptemp[el] = struct{}{}
	}
	for el := range maptemp {
		set = append(set, el)
	}
	fmt.Println(set)
}
