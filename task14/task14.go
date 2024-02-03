package main

import (
	"fmt"
	"reflect"
	"wb-l1/task11"
)

func print_type(thing interface{}) {
	t := reflect.TypeOf(thing)

	fmt.Println(t)
}

func main14() {
	print_type(1)
	print_type(1.1)
	print_type("1")
	print_type(task11.Set[int]{})
	print_type(make(<-chan chan task11.Set[int]))
}
