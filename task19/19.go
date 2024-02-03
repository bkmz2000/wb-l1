package main

import (
	"fmt"
	"strings"
)

func bad(s string) string { // wrong, does not recognise compound glyphs
	ret := strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Fprint(&ret, string(s[i]))
	}

	return ret.String()
}

func naiive(s string) string { // Simple, but very slow. O(n^2)?
	ret := ""

	for _, v := range s {
		ret = string(v) + ret
	}

	return ret
}

func good(s string) string { // O(n)
	rs := []rune(s)

	for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
		rs[l], rs[r] = rs[r], rs[l]
	}

	return string(rs)
}

func test_reverser(f func(string) string) {
	ss := []string{
		"Hellow world",
		"Äpfel und Bälle im öffentlichen!",
		"🌈 Hello, Universe! 🚀✨",
		"当太阳升起时，新的一天开始了。",
	}

	expected := []string{
		"dlrow wolleH",
		"!nehciltneffö mi elläB dnu lefpÄ",
		"✨🚀 !esrevinU ,olleH 🌈",
		"。了始开天一的新，时起升阳太当",
	}

	for i, s := range ss {
		if f(s) != expected[i] {
			fmt.Println("Expected", expected[i], "found", f(s))
		}
	}
}

func test_reversers() {
	fmt.Println("bad: ")
	test_reverser(bad) // works just with ASCII?

	fmt.Println("\nnaiive:")
	test_reverser(naiive) // ok

	fmt.Println("\ngood:")
	test_reverser(good) // ok
}
