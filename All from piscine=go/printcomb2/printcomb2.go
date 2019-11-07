package piscine

import "github.com/01-edu/z01"

func IntToRune(num int) (runenum rune) {
	switch num {
	case 0:
		runenum = '0'
	case 1:
		runenum = '1'
	case 2:
		runenum = '2'
	case 3:
		runenum = '3'
	case 4:
		runenum = '4'
	case 5:
		runenum = '5'
	case 6:
		runenum = '6'
	case 7:
		runenum = '7'
	case 8:
		runenum = '8'
	case 9:
		runenum = '9'
	}
	return runenum
}

func TwoDgNum(num int) {
	if num < 10 {
		z01.PrintRune('0')
		z01.PrintRune(IntToRune(num))
	} else {
		z01.PrintRune(IntToRune(num / 10))
		z01.PrintRune(IntToRune(num % 10))
	}
}

func PrintComb2() {

	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if i == 98 && j == 99 {
				TwoDgNum(i)
				z01.PrintRune(' ')
				TwoDgNum(j)
				z01.PrintRune('\n')
			} else {
				TwoDgNum(i)
				z01.PrintRune(' ')
				TwoDgNum(j)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
