package piscine

func Unmatch(arr []int) int {
	lengh := 0
	for i := range arr {
		lengh = i + 1
	}
	for index, num := range arr {
		for i := lengh - 1; i > index; i-- {
			if num == arr[i] {
				arr[index] = -1
				arr[i] = -1
			}
		}
	}
	for _, n := range arr {
		if n != -1 {
			return n
		}
	}
	return -1
}
