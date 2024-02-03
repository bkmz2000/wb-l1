package main

import (
	"fmt"
	"strings"
)

var justString string

func generateRandomString(length int) string {
	sb := strings.Builder{}

	symbols := "abcdefghijklmnopqrstuvwxyz123434567890"
	for i := 0; i < length; i++ {
		fmt.Fprint(&sb, symbols[i])
	}

	return sb.String()
}

func someFunc() {
	justString = generateRandomString(100)
}

func main() {
	someFunc()
}
