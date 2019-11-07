package piscine

func NRune(s string, n int) rune {
	strAsByte := []rune(s)
	var nrune rune
	for index, letter := range strAsByte {
		if index == n-1 {
			nrune = letter
		}
	}
	return nrune
}
