package piscine

func AlphaCount(str string) int {
	counter := 0
	word := []byte(str)
	for _, letter := range word {
		if (letter >= 65 && letter <= 90) ||
			(letter >= 97 && letter <= 122) {
			counter++
		}
	}
	return counter
}
