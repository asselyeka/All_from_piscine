package piscine

func Swap(a *int, b *int) {
	var buf int
	buf = *a
	*a = *b
	*b = buf
}
