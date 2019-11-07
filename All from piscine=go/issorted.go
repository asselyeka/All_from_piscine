package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	size := 0
	for i := range tab {
		size = i + 1
	}
	counterP := 0
	counterN := 0
	counter0 := 0
	if size > 1 {
		for i := 1; i < size; i++ {
			if f(tab[i], tab[i-1]) > 0 {
				counterP++
			} else if f(tab[i], tab[i-1]) < 0 {
				counterN++
			} else if f(tab[i], tab[i-1]) == 0 {
				counter0++
			}
		}
	}
	if counterP == size-1 || counterN == size-1 || counter0 == size-1 {
		return true
	}
	return false
}
