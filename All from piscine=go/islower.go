package piscine

func IsLower(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	for _, letter := range sAsRune {
		if checkAlplow(letter) {
			counter++
		}
	}
	if counter == StrLenAlplow(str) {
		return true
	}
	return false
}

func checkAlplow(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func StrLenAlplow(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
