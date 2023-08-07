package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(arr []int, value int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middlePoint := (left + right) / 2
		middleVal := arr[middlePoint]
		if middleVal == value {
			return middlePoint
		} else if value > middleVal {
			left = middlePoint + 1
		} else {
			right = middlePoint - 1
		}
	}
	return 0
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(arr, 8))
}
