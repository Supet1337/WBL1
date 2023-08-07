package main

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// воркер
func worker(wg *sync.WaitGroup, tasks <-chan int, id int) {
	for num := range tasks {
		fmt.Printf("Worker ID: %d: %d\n", id, num)
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	tasks := make(chan int)
	var N int
	fmt.Println("Count of workers:")
	fmt.Scan(&N)

	//Запуск воркеров
	wg.Add(N)
	for i := 0; i < N; i++ {
		go worker(&wg, tasks, i)
	}

	//Запись в канал
	go func() {
		j := 0
		for {
			j++
			tasks <- j
			time.Sleep(time.Second)
		}
	}()

	// Завершение программы
	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, syscall.SIGINT, syscall.SIGTERM)
	<-shutDown
	//Закрытие канала
	close(tasks)
	fmt.Println("Program was finished")
	wg.Wait()

}
