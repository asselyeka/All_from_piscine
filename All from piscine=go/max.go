package piscine

func Max(arr []int) int {
	size := 0
	for i := range arr {
		size = i + 1
	}
	if size > 0 {
		max := arr[0]
		for _, elem := range arr {
			if max < elem {
				max = elem
			}
		}
		return max
	}
	return 0
}
