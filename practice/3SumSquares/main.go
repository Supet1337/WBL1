package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

// Запись в канал квадарата
func square(el int, ch chan<- int) {
	ch <- el * el
}

func main() {
	mass := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)

	//Запуск горутин для подсчёта квадрата
	for _, el := range mass {
		go square(el, ch)
	}

	sum := 0
	//Чтение данных из канала
	for i := 0; i < len(mass); i++ {
		sum += <-ch
	}
	fmt.Println(sum)
	close(ch)

}
