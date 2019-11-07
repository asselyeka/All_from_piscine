package piscine

func IsUpper(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	for _, letter := range sAsRune {
		if checkAlpup(letter) {
			counter++
		}
	}
	if counter == StrLenAlpup(str) {
		return true
	}
	return false
}

func checkAlpup(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func StrLenAlpup(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
