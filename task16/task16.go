package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Ordered interface {
	LessThen(other Ordered) bool
}

type Int int

func (this Int) LessThen(other Ordered) bool {
	return this < other.(Int)
}

func swap[T Ordered](arr []T, l int, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func partition[T Ordered](arr []T, l int, r int) int {
	pivot := arr[(l+r)/2]

	i, j := l, r

	for i <= j {
		for arr[i].LessThen(pivot) {
			i++
		}

		for pivot.LessThen(arr[j]) {
			j--
		}

		if i <= j {
			swap(arr, i, j)
			i++
			j--
		}
	}

	return i
}

func qsortlr[T Ordered](arr []T, l int, r int) {
	if l >= r {
		return
	}

	q := partition(arr, l, r)

	qsortlr(arr, l, q-1)
	qsortlr(arr, q, r)
}

func qsort[T Ordered](arr []T) {
	qsortlr(arr, 0, len(arr)-1)
}

func test_qsort() {
	for test := 0; test < 100; test++ {
		length := rand.Intn(1000) + 5
		a := make([]int, length)
		b := make([]Int, length)

		for i := 0; i < length; i++ {
			new := rand.Intn(length)
			a[i] = new
			b[i] = Int(new)
		}

		sort.Ints(a)
		qsort(b)

		for i := 0; i < length; i++ {
			if Int(a[i]) != b[i] {
				panic(fmt.Sprint(a, b))
			}
		}
	}

	fmt.Println("Ok")
}
