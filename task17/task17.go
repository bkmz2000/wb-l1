package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binary_search(arr []int, n int) int {
	l, r := 0, len(arr)

	for l < r {
		m := (l + r) / 2

		if arr[m] == n {
			return m
		}

		if arr[m] > n {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func test_binary_search() {
	for test := 0; test < 1000; test++ {
		length := rand.Intn(10) + 5
		a := make([]int, length)

		for i := 0; i < length; i++ {
			a[i] = rand.Intn(length)
		}

		n := rand.Intn(1000)
		a[rand.Intn(len(a))] = n
		sort.Ints(a)

		ind := binary_search(a, n)

		if a[ind] != n {
			panic("Wrong")
		} else {
			fmt.Println("Ok", test)

		}
	}

	fmt.Println("Ok")
}
