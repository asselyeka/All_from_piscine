package piscine

func ByteToInt(num byte) int {
	var runenum int
	switch num {
	case 48:
		runenum = 0
	case 49:
		runenum = 1
	case 50:
		runenum = 2
	case 51:
		runenum = 3
	case 52:
		runenum = 4
	case 53:
		runenum = 5
	case 54:
		runenum = 6
	case 55:
		runenum = 7
	case 56:
		runenum = 8
	case 57:
		runenum = 9
	}
	return runenum
}

func StrLen(str string) int {
	var count int
	var a int
	for index, word := range str {
		if word == 233 {
			a = -1
		}
		count = index + 1 + a
	}
	return count
}

func BasicAtoi(s string) int {
	lenth := StrLen(s)
	var num int = 0
	str := []byte(s)
	for i := 0; i < lenth; i++ {
		if str[i] > 48 && str[i] < 58 {
			for j := i; j < lenth; j++ {
				num = num*10 + ByteToInt(s[j])
			}
			break
		}
	}
	return num
}
