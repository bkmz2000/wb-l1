package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sender(lifetime time.Duration, c chan int, wg *sync.WaitGroup) {
	start := time.Now()

	endTime := time.Now().Add(lifetime)

	for time.Now().Before(endTime) {
		c <- rand.Intn(100)
	}

	close(c)
	wg.Done()
	fmt.Println("Sender worked for ", time.Now().Sub(start).Milliseconds(), "ms")
}

func printer(c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Println(n)
	}
	wg.Done()
}

func main() {
	c := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(2, c, &wg)
	go printer(c, &wg)

	wg.Wait()
}
