package piscine

func SortWordArr(array []string) {
	size := 0
	for i := range array {
		size = i + 1
	}
	var temp string
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
