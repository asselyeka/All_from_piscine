package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lengthOfArg := 0
	for i := range arg {
		lengthOfArg = i + 1
	}

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(str string) {
	arrayStr := []rune(str)
	lenStr := 0
	for i := range arrayStr {
		lenStr = i + 1
	}

	for i := 0; i < lenStr; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	}
	return false
}
