package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

import (
	"fmt"
)

var justString string

func createHugeString(size int) (str string) {
	for i := 0; i < size; i++ {
		str += string(rune(size))
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	// 1 << 100 Это 100 байт
	// Но мы можем положить туда символ больше 1 байта
	justString = v[:100]
}

func someGoodFunc() {
	v := createHugeString(1 << 10)
	// конвертируем строку в []rune, чтобы использовать тяжеловесные символы
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
	fmt.Println(justString) // 50 символов(( Но мы же хотели 100
	someGoodFunc()
	fmt.Println(justString) // А тут уже 100))
}
