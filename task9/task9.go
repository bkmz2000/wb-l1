package main

import (
	"fmt"
	"math/rand"
)

func writer(arr []int, ci chan int) {
	for n := range arr {
		ci <- n
	}

	close(ci)
}

func squarer(ci chan int, co chan int) {
	for n := range ci {
		co <- n * n
	}

	close(co)
}

func main() {
	length := rand.Intn(1000)

	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(1000)
	}

	ci := make(chan int, 10)
	co := make(chan int, 10)

	go writer(arr, ci)
	go squarer(ci, co)

	for n := range co {
		fmt.Println(n)
	}
}
