package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
type Counter struct {
	i int
}

func worker(wg *sync.WaitGroup, m *sync.Mutex, c *Counter, i int) {
	m.Lock()
	c.i++
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	c := Counter{0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&wg, &m, &c, i)
	}

	wg.Wait()
	fmt.Println(c.i)
}
