package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	size := 0
	var output []rune
	for {
		symbol, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, symbol)
		size++
	}

	if size > 0 {
		if output[0] != 'o' && output[0] != '/' && output[0] != 'A' {
			fmt.Println("Not a Raid function")
			return
		}
		if



		InputFunc(output)
	}
}

func coord(output []rune) (x, y int) {
	X := 0
	Y := 0
	flag := true
	for _, symbol := range output {
		if symbol == '\n' {
			Y++
			flag = false
		} else {
			if flag == true {
				X++
			}
		}
	}
	return X, Y
}

func InputFunc(output []rune) {
	x, y := coord(output)
	leftdownSym := ((x * y) + y - 1) - x
	rightdownSym := (x * y) + y - 2

	if output[0] == 'o' {
		fmt.Printf("[raid1a] [%d] [%d]", x, y)
	} else if output[0] == '/' {
		fmt.Printf("[raid1b] [%d] [%d]", x, y)
	} else if output[0] == 'A' {
		if x == 1 && y == 1 {
			fmt.Printf("[raid1c] [%d] [%d] || [raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y, x, y)
		} else if x == 1 && y > 0 && output[leftdownSym] == 'C' {
			fmt.Printf("[raid1c] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		} else if y == 1 && x > 0 && output[x-1] == 'C' {
			fmt.Printf("[raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		} else if output[leftdownSym] == 'C' && output[rightdownSym] == 'C' {
			fmt.Printf("[raid1c] [%d] [%d]", x, y)
		} else if output[leftdownSym] == 'A' && output[rightdownSym] == 'C' {
			fmt.Printf("[raid1d] [%d] [%d]", x, y)
		} else if output[leftdownSym] == 'C' && output[rightdownSym] == 'A' {
			fmt.Printf("[raid1e] [%d] [%d]", x, y)
		}
	}
	fmt.Println()
}
