package piscine

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

func SortIntegerTable(table []int) {
	var size int
	var temp int
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
}
