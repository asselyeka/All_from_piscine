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

func PrintComb() {

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				if i == 7 && j == 8 && k == 9 {
					z01.PrintRune(IntToRune(i))
					z01.PrintRune(IntToRune(j))
					z01.PrintRune(IntToRune(k))
					z01.PrintRune('\n')
					break
				} else {
					z01.PrintRune(IntToRune(i))
					z01.PrintRune(IntToRune(j))
					z01.PrintRune(IntToRune(k))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

}
