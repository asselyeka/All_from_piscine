package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	l := LenLen(base)
	if CheckBase(base) == false || l < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else if nbr > 2000000000 || nbr < -2000000000 {
		PrintNbr(nbr)
	} else {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr *= -1
		}
		str := ""
		for nbr >= l {
			str += string(base[nbr%l])
			nbr = nbr / l
		}
		str += string(base[nbr])
		str = RevStr(str)
		for _, letter := range str {
			z01.PrintRune(letter)
		}
	}
}
func CheckBase(base string) bool {
	for index, letter := range base {
		for i := index + 1; i < LenLen(base); i++ {
			if letter == '-' || letter == '+' || letter == rune(base[i]) {
				return false
			}
		}
	}
	return true
}
func RevStr(str string) string {
	var rev string
	l := LenLen(str)
	for i := l - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

func LenLen(str string) int {
	count := 0
	a := []rune(str)
	for index := range a {
		count = index + 1
	}
	return count
}
