package piscine

func LastRune(s string) rune {
	strAsRune := []rune(s)
	var nrune rune
	for _, letter := range strAsRune {
		nrune = letter
	}
	return nrune
}
