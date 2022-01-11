package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func print(conv func(string) string) func(string) {
	return func(s string) {
		fmt.Printf("(%[1]q, %+[1]q) -> (%[2]q, %+[2]q)\n", s, conv(s))
	}
}

// 正準等価性
func canonicalEquivalence() {
	fmt.Println("canonical equivalence")

	f := print(norm.NFD.String)

	f("a")
	f("é")
	f("Å")
	f("①")
}

// 互換等価性
func compatibilityEquivalence() {
	fmt.Println("compatibility equivalence")

	f := print(norm.NFKD.String)

	f("a")
	f("é")
	f("①")
	f("Ⅻ")
}

func main() {
	canonicalEquivalence()
	fmt.Println()
	compatibilityEquivalence()
}
