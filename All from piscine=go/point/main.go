package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	var num int
	var n int
	var revnum int
	points := point{}
	setPoint(&points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	num = points.x
	revnum = 0
	for num != 0 {
		n = num % 10
		revnum = revnum*10 + n
		num /= 10
	}
	for revnum != 0 {
		n = revnum % 10
		counter := 0
		for i := '0'; i <= '9'; i++ {
			if counter == n {
				z01.PrintRune(i)
				break
			} else {
				counter++
			}
		}
		revnum /= 10
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	num = points.y
	revnum = 0
	for num != 0 {
		n = num % 10
		revnum = revnum*10 + n
		num /= 10
	}
	for revnum != 0 {
		n = revnum % 10
		counter := 0
		for i := '0'; i <= '9'; i++ {
			if counter == n {
				z01.PrintRune(i)
				break
			} else {
				counter++
			}
		}
		revnum /= 10
	}
	z01.PrintRune('\n')
}
