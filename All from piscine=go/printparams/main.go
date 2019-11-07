package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	len := 0
	for index := range arguments {
		len = index + 1
	}
	for i := 1; i < len; i++ {
		strAsRune := []rune(arguments[i])
		for _, letter := range strAsRune {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
