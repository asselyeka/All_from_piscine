package piscine

func Join(strs []string, sep string) string {
	var str string
	for index, words := range strs {
		if index != 0 {
			str += sep

		}
		str += words
	}
	return str
}
