package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

func main() {
	str := "abcd"
	str = strings.ToLower(str)
	fmt.Println(checkunique(str))
}

func checkunique(str string) bool {
	mp := make(map[rune]struct{})
	for _, el := range str {
		mp[el] = struct{}{}
	}
	if len(mp) == len(str) {
		return true
	}
	return false
}
