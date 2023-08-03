package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func worker(wg *sync.WaitGroup, m *sync.Mutex, mp *map[int]int, i int) {
	m.Lock()
	(*mp)[i] = i
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	mp := make(map[int]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&wg, &m, &mp, i)
	}

	wg.Wait()
	fmt.Println(mp)
}
