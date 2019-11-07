package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, str := range tab {
		if f(str) {
			counter++
		}
	}
	return counter
}
