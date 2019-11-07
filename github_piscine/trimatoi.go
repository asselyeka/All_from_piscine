package piscine

func TrimAtoi(s string) int {
	result := 0
	strAsByte := []byte(s)
	isNegative := false
	var numAsStr string
	for index, symbol := range strAsByte {
		// '0' = 48 '-' = 45
		if symbol > 48 && symbol < 58 {
			for i := 0; i < index; i++ {
				if strAsByte[i] == 45 {
					isNegative = true
				}
			}
			for i := index; i < StrLen(s); i++ {
				if strAsByte[i] > 47 && strAsByte[i] < 58 {
					numAsStr += string(strAsByte[i])
				}
			}
			break
		}
	}
	numAsByte := []byte(numAsStr)
	for _, symbol := range numAsByte {
		result = result*10 + ByteToInt(symbol)
	}
	if isNegative == true {
		result *= -1
	}
	return result
}

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
	for index := range str {
		count = index + 1
	}
	return count
}
