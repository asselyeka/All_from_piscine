package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		if n < 0 {
			z01.PrintRune('-')
		}
		str := SortStr(Itoa(n))
		for _, letter := range str {
			z01.PrintRune(letter)
		}
	}
}

func Itor(n int) string {
	switch n {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	default:
		return "9"
	}
}

func Itoa(n int) string {
	size := 0
	numAsStr := ""
	num := n
	for num != 0 {
		num /= 10
		size++
	}
	for i := 0; i < size; i++ {
		numAsStr += Itor(n % 10)
		n /= 10
	}
	return numAsStr
}

func SortStr(numAsStr string) string {
	table := []byte(numAsStr)
	size := 0
	var temp byte
	for index := range table {
		size = index + 1
	}
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
	return string(table)
}
