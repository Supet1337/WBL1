package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	var str = "snow dog sun"
	reverse := strings.Split(str, " ")
	for i := 0; i < len(reverse)/2; i++ {
		reverse[i], reverse[len(reverse)-1-i] = reverse[len(reverse)-1-i], reverse[i]
	}

	fmt.Println(reverse)
}
