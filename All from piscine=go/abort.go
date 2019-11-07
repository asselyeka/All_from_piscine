package piscine

func Abort(a, b, c, d, e int) int {
	var temp int
	if a > b {
		temp = a
		a = b
		b = temp
	}
	if b > c {
		if a > c {
			temp = a
			a = c
			c = b
			b = temp
		} else {
			temp = b
			b = c
			c = temp
		}
	}
	if b > d {
		if a > d {
			temp = d
			d = c
			c = b
			b = a
			a = temp
		} else {
			temp = d
			d = c
			c = b
			b = temp
		}
	} else {
		if c > d {
			temp = d
			d = c
			c = temp
		}
	}
	if c > e {
		if b > e {
			c = b
		} else {
			c = e
		}
	}
	return c
}
