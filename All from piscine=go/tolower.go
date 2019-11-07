package piscine

func ToLower(s string) string {
	sAsByte := []byte(s)
	for index, letter := range sAsByte {
		if letter <= 90 && letter >= 65 {
			sAsByte[index] = letter + 32
		}
	}
	return string(sAsByte)
}
