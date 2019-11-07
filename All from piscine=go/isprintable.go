package piscine

func IsPrintable(str string) bool {
	sAsRune := []rune(str)
	for _, letter := range sAsRune {
		if checkPrint(letter) {
			return false
		}
	}
	return true
}

func checkPrint(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
