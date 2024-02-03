package main

import (
	"fmt"
	"strings"
)

func reverse_words(s string) string {
	words := strings.Fields(s)
	sb := strings.Builder{}

	for i := range words {
		fmt.Fprintf(&sb, "%s ", words[len(words)-1-i])
	}

	return sb.String()
}

func main() {
	fmt.Println(reverse_words("snow dog sun"))
}
