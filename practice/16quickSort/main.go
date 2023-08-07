package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int

	for _, el := range arr[1:] {
		if el <= pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}
	result := append(quickSort(left), pivot)
	result = append(result, quickSort(right)...)
	return result
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	arr = quickSort(arr)
	fmt.Println(arr)
}
