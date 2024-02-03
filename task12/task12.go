package main

import (
	"fmt"
	"wb-l1/task11"
)

func main12() {
	set := task11.NewSet[string]()

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, s := range data {
		set.Add(s)
	}

	for s := range set.Iterator() {
		fmt.Println(s)
	}
}
