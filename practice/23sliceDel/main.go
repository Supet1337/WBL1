package main

import (
	"errors"
	"fmt"
	"log"
)

//Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr, err := sliceDel(arr, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
}
func sliceDel(arr []int, idx int) ([]int, error) {
	if idx >= len(arr) || idx < 0 {
		return nil, errors.New("Bad index")
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr, nil
}
