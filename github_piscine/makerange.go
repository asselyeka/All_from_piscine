package piscine

func MakeRange(min, max int) []int {
	if min < max {
		answer := make([]int, max-min)
		for index := range answer {
			answer[index] = min + index
		}
		return answer
	}
	return nil
}
