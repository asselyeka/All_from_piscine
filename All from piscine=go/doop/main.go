package main

import (
	"fmt"
	"os"
)

func Div(a, b int) int {
	return a / b
}
func Mult(a, b int) int {
	return a * b
}
func Mod(a, b int) int {
	return a % b
}
func Plus(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

type Calculator struct {
	symbol  rune
	arrFunc func(int, int) int
}

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	lenArg := 0
	arrCalculator := []Calculator{
		{'/', Div},
		{'*', Mult},
		{'%', Mod},
		{'+', Plus},
		{'-', Minus},
	}
	for i := range arguments {
		lenArg = i + 1
	}
	result := 0
	errSym := -1
	if lenArg == 3 {
		if IsNumericD(arguments[0]) && IsNumericD(arguments[2]) {
			for i, fun := range arrCalculator {
				if arguments[1] == string(fun.symbol) {
					errSym++
					if AtoiD(arguments[2]) == 0 && i == 0 {
						fmt.Println("No division by 0")
						break
					} else if AtoiD(arguments[2]) == 0 && i == 2 {
						fmt.Println("No Modulo by 0")
						break
					} else {
						result = fun.arrFunc(AtoiD(arguments[0]), AtoiD(arguments[2]))
						fmt.Println(result)
					}
				}
			}
			if errSym == -1 {
				fmt.Println(0)
			}
		} else {
			fmt.Println(0)
		}

	}

}

func IsNumericD(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	for _, letter := range sAsRune {
		if checkNumD(letter) {
			counter++
		}
	}
	if sAsRune[0] == '-' {
		counter++
	}
	if counter == StrLenNumD(str) {
		return true
	}
	return false
}

func checkNumD(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func StrLenNumD(str string) int {
	var count int
	strAsByte := []rune(str)
	for index := range strAsByte {
		count = index + 1
	}
	return count
}

func BToIntD(num byte) int {
	var runenum int
	switch num {
	case 48:
		runenum = 0
	case 49:
		runenum = 1
	case 50:
		runenum = 2
	case 51:
		runenum = 3
	case 52:
		runenum = 4
	case 53:
		runenum = 5
	case 54:
		runenum = 6
	case 55:
		runenum = 7
	case 56:
		runenum = 8
	case 57:
		runenum = 9
	}
	return runenum
}

func SLenD(str string) int {
	var count int
	var a int
	for index, word := range str {
		if word == 233 {
			a = -1
		}
		count = index + 1 + a
	}
	return count
}

func BasicAtoi2D(s string, ind int) int {
	lenth := SLenD(s)
	var num int = 0
	str := []byte(s)
	for i := ind; i < lenth; i++ {
		if str[i] > 47 && str[i] < 58 {
			for j := i; j < lenth; j++ {
				if str[j] > 47 && str[j] < 58 {
					num = num*10 + BToIntD(s[j])
				} else {
					return 0
				}
			}
			break
		} else {
			break
		}

	}
	return num
}

func AtoiD(s string) int {
	str := []byte(s)
	var num int
	if SLenD(s) >= 2 {

		if str[0] == 43 {
			num = BasicAtoi2D(s, 1)
		} else if str[0] == 45 {
			num = (-1) * BasicAtoi2D(s, 1)
		} else {
			num = BasicAtoi2D(s, 0)
		}
	} else {
		num = BasicAtoi2D(s, 0)
	}

	return num
}
