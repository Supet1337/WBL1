package main

/*Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/
import (
	"fmt"
	"sync"
)

func main() {
	mass := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, el := range mass {
		wg.Add(1)
		go func(el int) {
			fmt.Println(el * el)
			defer wg.Done()
		}(el)
	}
	wg.Wait()
}
