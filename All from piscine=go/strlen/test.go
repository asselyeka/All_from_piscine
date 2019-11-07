package main

import (
	"fmt"

	piscine ".."
)

func main() {
	str := "Héllo!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
