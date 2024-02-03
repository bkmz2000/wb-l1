package main

import (
	"fmt"
	"sync"
	"time"
)

func sq_from_chan(from chan int, to chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range from {
		to <- n * n
		fmt.Println(n, num)
	}
}

func concurent_sq_sum(nums []int) {
	start := time.Now()

	from := make(chan int, 100)
	to := make(chan int, 100)

	var wg sync.WaitGroup

	for i := 0; i < len(nums); i++ {
		from <- nums[i]
	}

	close(from)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sq_from_chan(from, to, i, &wg)
	}

	go func() {
		wg.Wait()
		close(to)
	}()

	s := 0

	for n := range to {
		s += n
	}

	fmt.Println(s, time.Now().Sub(start))
}
