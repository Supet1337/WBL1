# Чем отличаются RWMutex от Mutex?


### RWMutex отличается от обычного Mutex тем, что первый позволяет заблокировать объект на чтение\запись  
### RWMutex подходит для ситуаций, когда большая часть операций является чтением, что позволяет параллельно считывать данные несольким горутинам эффективно
```
package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwMutex sync.RWMutex
	counter := 0

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 50; i++ {

		go func() {
			defer wg.Done()
			rwMutex.RLock() //Позволяет нескольким горутинам читать данные
			_ = counter
			rwMutex.RUnlock()
		}()

		go func() {
			defer wg.Done()
			rwMutex.Lock() //Блокировка на чтение
			counter++
			rwMutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}


```
