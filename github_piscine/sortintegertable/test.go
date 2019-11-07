package main

import (
	"fmt"

	piscine ".."
)

func main() {
	s := []int{12, 10, 7, 6, 5, 4, 3, 2, 1, 0}
	a := []int{5, -18, 23, -96}
	b := []int{57, -71, -89, -86}
	c := []int{-93, -99, 93, 74}
	piscine.SortIntegerTable(s)
	piscine.SortIntegerTable(a)
	piscine.SortIntegerTable(b)
	piscine.SortIntegerTable(c)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// sortintegertable_test.go:27: SortIntegerTable([5 -18 23 -96]), table == [-18 5 23 -96] instead of [-96 -18 5 23]
// sortintegertable_test.go:27: SortIntegerTable([57 -71 -89 -86]), table == [-89 -71 57 -86] instead of [-89 -86 -71 57]
// sortintegertable_test.go:27: SortIntegerTable([-93 -99 93 74]), table == [-99 -93 93 74] instead of [-99 -93 74 93]
