package piscine

func ForEach(f func(int), arr []int) {
	for _, n := range arr {
		f(n)
	}
}
