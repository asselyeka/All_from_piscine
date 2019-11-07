package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, word := range args {
		if i == 0 {
			str = word
		} else {
			str += "\n" + word
		}
	}
	return str
}
