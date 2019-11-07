package piscine

func IsNumeric(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	for _, letter := range sAsRune {
		if checkNum(letter) {
			counter++
		}
	}
	if counter == StrLenNum(str) {
		return true
	}
	return false
}

func checkNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func StrLenNum(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
