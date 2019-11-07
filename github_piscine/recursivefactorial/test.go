package main

import (
	"fmt"

	piscine ".."
)

func main() {
	arg := 5
	fmt.Println(piscine.RecursiveFactorial(arg))
	arg1 := -4
	fmt.Println(piscine.RecursiveFactorial(arg1))
	arg2 := 0
	fmt.Println(piscine.RecursiveFactorial(arg2))
}
