package piscine

func UltimateDivMod(a *int, b *int) {
	var dev int
	var mod int
	dev = *a / (*b)
	mod = *a % (*b)
	*a = dev
	*b = mod
}
