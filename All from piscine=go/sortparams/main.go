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

	table := []string(arguments[1:])
	size := len - 1
	var temp string
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if table[j] > table[j+1] {
				// меняем элементы местами
				temp = table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}

	for i := 0; i < len-1; i++ {
		for _, letter := range table[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
