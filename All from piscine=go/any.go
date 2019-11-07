package piscine

func Any(f func(string) bool, arr []string) bool {
	counter := 0
	for _, str := range arr {
		if f(str) {
			counter++
		}
	}
	if counter > 0 {
		return true
	}
	return false
}
