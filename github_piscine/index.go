package piscine

func Index(s string, toFind string) int {
	strAsByte := []rune(s)
	toFindByte := []rune(toFind)
	indexFind := -1
	if toFind == "" {
		indexFind = 0
	} else {
		for index, letter := range strAsByte {
			if letter == toFindByte[0] {
				j := 0
				count := 0
				if index+StrLenn(toFind) <= StrLenn(s) {
					for i := index; i < StrLenn(toFind); i++ {
						if strAsByte[i] == toFindByte[j] {
							count++
						}
						j++
					}
					if count == j {
						indexFind = index
						break
					}
				}
			}
		}
	}
	return indexFind
}

func StrLenn(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}
