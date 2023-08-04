package main

import "fmt"

// Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chx := make(chan int, len(arr))
	chx2 := make(chan int, len(arr))
	wrtCh(arr, chx, chx2)
	prtCh(chx2)

}

func wrtCh(arr []int, chx chan int, chx2 chan int) {
	for _, el := range arr {
		chx <- el
		chx2 <- el * el
	}
	close(chx)
	close(chx2)
}

func prtCh(chx2 chan int) {
	for el := range chx2 {
		fmt.Println(el)
	}

}
