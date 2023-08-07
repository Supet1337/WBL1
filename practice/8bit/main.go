package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var num int64 = 63
	fmt.Printf("old  %064b\n", num)
	changebit(num, 1, 0)
}

func changebit(num int64, pos, bit int64) {
	if bit == 1 {
		mask := (bit << pos)
		fmt.Printf("pos %d\n", pos)
		fmt.Printf("mask %064b\n", mask)
		num = mask | num

	} else if bit == 0 {
		mask := (int64(1) << pos)
		fmt.Printf("pos %d\n", pos)
		fmt.Printf("mask %064b\n", mask)
		num = ^mask & num
	}
	fmt.Printf("new  %064b\n", num)
}
