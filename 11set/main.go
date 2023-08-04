package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	firstSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 34, 23, 2121, 1, 7}
	secondSet := []int{1, 2, 3, 4, 5, 6, 11, 12, 14, 7}
	resultMap := make(map[int]struct{})
	resultSet := []int{}

	for _, el := range firstSet {
		resultMap[el] = struct{}{}
	}

	for _, el := range secondSet {
		if _, ok := resultMap[el]; ok {
			resultSet = append(resultSet, el)
		}
	}

	fmt.Println(resultSet)
}
