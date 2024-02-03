package main

import (
	"fmt"
	"sync"
)

func worker(ci chan int, co chan int) {
	for n := range ci {
		co <- n * n
	}
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	ci := make(chan int, 5)
	co := make(chan int, 5)

	for _, n := range arr {
		ci <- n
	}

	close(ci)

	wg := sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			worker(ci, co)
			wg.Done()
		}()
	}

	s := 0

	go func() {
		wg.Wait()
		close(co)
	}()

	for n := range co {
		s += n
	}

	fmt.Print(s)
}
