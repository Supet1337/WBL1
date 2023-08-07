package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {
	var str = "главрыба"
	runeStr := []rune(str)

	for i := 0; i < len(runeStr)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-1-i] = runeStr[len(runeStr)-1-i], runeStr[i]
	}

	fmt.Println(string(runeStr))
}
