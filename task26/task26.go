package main

import "fmt"

func all_unique(s string) bool {
	counts := make(map[rune]int)

	for _, r := range s {
		_, ok := counts[r]
		if ok {
			return false
		}

		counts[r] = 1
	}

	return true
}

func main() {
	ss := [...]string{
		"Hellow world",
		"Äpfel und Äpfel",
		"world",
		"Äpfel",
		"🚀✨🌈",
		"🚀✨🌈✨",
		"当太阳",
		"当太阳太",
	}

	for _, s := range ss {
		fmt.Println(all_unique(s))
	}
}
