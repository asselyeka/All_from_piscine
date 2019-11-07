package piscine

func Map(f func(int) bool, arr []int) []bool {
	lenA := 0
	for i := range arr {
		lenA = i + 1
	}
	arrBool := make([]bool, lenA)
	for i, n := range arr {
		arrBool[i] = f(n)
	}
	return arrBool
}
