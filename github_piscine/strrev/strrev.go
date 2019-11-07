package piscine

func StrRev(s string) string {
	var count int
	var a int
	for index, word := range s {
		if word == 233 {
			a = -1
		}
		count = index + a
	}
	str := s
	revstr := []byte(s)
	j := 0
	for i := count; i >= 0; i-- {
		revstr[i] = byte(s[j])
		j++
	}
	finalstr := string(revstr)

	return finalstr
}
