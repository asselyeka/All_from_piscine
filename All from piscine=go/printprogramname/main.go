package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[0]
	nameAsRune := []rune(argument)
	for _, letter := range nameAsRune {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
