package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, str := range table {
		strAsRune := []rune(str)
		for _, l := range strAsRune {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
