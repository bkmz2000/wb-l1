package main

import "fmt"

func main13() {
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)
}
