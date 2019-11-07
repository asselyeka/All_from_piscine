package piscine

func FirstRune(s string) rune {
	var fistRune rune
	for index, letter := range s {
		if index == 0 {
			fistRune = letter
		}
	}
	return fistRune
}
