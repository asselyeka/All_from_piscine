package piscine

func ToUpper(s string) string {
	sAsByte := []byte(s)
	for index, letter := range sAsByte {
		if letter <= 122 && letter >= 97 {
			sAsByte[index] = letter - 32
		}
	}
	return string(sAsByte)
}
