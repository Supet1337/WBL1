package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func gorutine1(stop <-chan bool) {
	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("gorutine1 working")
			time.Sleep(1 * time.Second)
		case <-stop:
			// Получен сигнал об остановке
			fmt.Println("gorutine1 done")
			return
		}
	}
}

func gorutine2(c chan<- int) {
	for i := 0; i <= 3; i++ {
		c <- i
	}
	close(c) // close channel
	fmt.Println("gorutine2 done")
}

func gorutine3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("gorutine3 working")
	time.Sleep(2 * time.Second)
}

func gorutine4() {
	fmt.Println("gorutine4 working")
	time.Sleep(2 * time.Second)
	return
}

func gorutine5() {
	fmt.Println("gorutine5 working")
	time.Sleep(2 * time.Second)
	runtime.Goexit()
}

func gorutine6(ctx context.Context) {
	fmt.Println("gorutine6 working")
	time.Sleep(2 * time.Second)
	<-ctx.Done()
}

func main() {
	//Создаём канал для 1 способа
	ch1 := make(chan bool)
	go gorutine1(ch1)

	//отправляем значение в канал
	time.Sleep(1 * time.Second)
	ch1 <- true

	// Ждем завершения горутины
	time.Sleep(1 * time.Second)

	//------------------------------------
	//2 способ, используя цикл for и закрытие канала
	ch2 := make(chan int)
	go gorutine2(ch2)

	for val := range ch2 {
		time.Sleep(1 * time.Second)
		fmt.Printf("gorutine2 working, value=%d\n", val)
	}
	time.Sleep(1 * time.Second)

	//------------------------------------
	//3 способ, waitgroup

	wg := sync.WaitGroup{}
	wg.Add(1)
	go gorutine3(&wg)

	wg.Wait()
	fmt.Println("gorutine3 done")

	//------------------------------------
	//4 способ return
	go gorutine4()
	time.Sleep(1 * time.Second)
	fmt.Println("gorutine4 done")

	//------------------------------------
	//5 способ runtime.Goexit()
	go gorutine5()
	time.Sleep(1 * time.Second)
	fmt.Println("gorutine5 done")

	//------------------------------------
	//6 способ использование контекста
	ctx := context.Background()
	go gorutine6(ctx)
	time.Sleep(1 * time.Second)
	fmt.Println("gorutine6 done")

}
