package piscine

func Compact(ptr *[]string) int {
	counter := 0
	for _, arr := range *ptr {
		if arr != "" {
			counter++
		}
	}
	finalPtr := make([]string, counter)
	i := 0
	for _, arr := range *ptr {
		if arr != "" {
			finalPtr[i] = arr
			i++
		}
	}
	*ptr = finalPtr
	return counter
}
