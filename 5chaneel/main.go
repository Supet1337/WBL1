package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// читает данные из канала и выводит их
func worker(wg *sync.WaitGroup, ch <-chan int) {
	for num := range ch {
		fmt.Printf("Recive: %d\n", num)
	}
	defer wg.Done()
}

func main() {

	wg := sync.WaitGroup{}

	var n string
	fmt.Println("Time")
	fmt.Scan(&n)
	t, err := time.ParseDuration(n + "s")
	if err != nil {
		panic(err)
	}

	input := make(chan int)
	wg.Add(1)
	go worker(&wg, input)

	go func() {
		timer := time.NewTimer(t)
		defer timer.Stop()
		for {
			select {
			default:
				a := rand.Int()
				time.Sleep(1 * time.Second)
				input <- a
				fmt.Printf("Send: %d\n", a)
			case <-timer.C:
				fmt.Println("times up")
				close(input)
				return
			}
		}
	}()

	wg.Wait()
}
