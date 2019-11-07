package piscine

func BasicJoin(strs []string) string {
	var str string
	for _, words := range strs {
		str += words
	}
	return str
}
