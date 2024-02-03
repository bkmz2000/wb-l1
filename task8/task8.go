package main

import "fmt"

func set_bit(n int64, i int) int64 {
	return n | (1 << i)
}

func main() {
	n := set_bit(0, 10)
	fmt.Printf("%b \n", n)
}
