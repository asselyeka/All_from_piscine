package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
