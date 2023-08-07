package main

import (
	"fmt"
	"time"
)

//разработать собственную функцию sleep.

func main() {
	fmt.Println("Sleep")
	sleep(time.Second * 5)
	fmt.Println("Wake up")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
