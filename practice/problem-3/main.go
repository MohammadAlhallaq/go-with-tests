package main

import (
	"fmt"
	"time"
)

func slowTask(name string, d time.Duration, ch chan string) {

	time.Sleep(d)
	ch <- name

}

func main() {
	ch := make(chan string)

	go slowTask("mohmmad", time.Second*3, ch)
	go slowTask("wow", time.Second*1, ch)
	go slowTask("ewww", time.Second*2, ch)

	for i := 0; i < 3; i++ {
		taskName := <-ch
		fmt.Println("done:", taskName)
	}

}
