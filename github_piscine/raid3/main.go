package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	output := getOutput()
	InputFunc(output)
}

func getOutput() []rune {
	input, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	if (input.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("Not a Raid function")
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)

	}
	if output == nil {
		fmt.Println("Not a Raid function")
		os.Exit(0)
	}
	return output
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
	raid1aStr := []rune(Raid1aStr(x, y))
	raid1bStr := []rune(Raid1bStr(x, y))
	raid1cStr := []rune(Raid1cStr(x, y))
	raid1dStr := []rune(Raid1dStr(x, y))
	raid1eStr := []rune(Raid1eStr(x, y))

	raid1a := true
	raid1b := true
	raid1c := true
	raid1d := true
	raid1e := true

	size := 0
	for i := range output {
		size = i + 1

	}
	for i := 0; i < size; i++ {
		if output[i] != raid1aStr[i] {
			raid1a = false
		}
	}
	for i := 0; i < size; i++ {
		if output[i] != raid1bStr[i] {
			raid1b = false
		}
	}
	for i := 0; i < size; i++ {
		if output[i] != raid1cStr[i] {
			raid1c = false
		}
	}
	for i := 0; i < size; i++ {
		if output[i] != raid1dStr[i] {
			raid1d = false
		}
	}
	for i := 0; i < size; i++ {
		if output[i] != raid1eStr[i] {
			raid1e = false
		}
	}

	if raid1a {
		fmt.Printf("[raid1a] [%d] [%d]", x, y)
	} else if raid1b {
		fmt.Printf("[raid1b] [%d] [%d]", x, y)
	} else if raid1c || raid1d || raid1e {
		if raid1c && raid1d && raid1e {
			fmt.Printf("[raid1c] [%d] [%d] || [raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y, x, y)
		} else if raid1c && raid1e {
			fmt.Printf("[raid1c] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		} else if raid1d && raid1e {
			fmt.Printf("[raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		} else if raid1c {
			fmt.Printf("[raid1c] [%d] [%d]", x, y)
		} else if raid1d {
			fmt.Printf("[raid1d] [%d] [%d]", x, y)
		} else if raid1e {
			fmt.Printf("[raid1e] [%d] [%d]", x, y)
		}
	} else {
		fmt.Println("Not a Raid function")
	}
	fmt.Println()
}

func Raid1aStr(x, y int) string {
	str := ""
	x--
	y--
	if x >= 0 && y >= 0 {
		for i := 0; i <= y; i++ {
			for j := 0; j <= x; j++ {
				if i == 0 && j == 0 ||
					i == y && j == x ||
					i == y && j == 0 ||
					i == 0 && j == x {
					str += "o"
					//	z01.PrintRune('')
				} else if i == 0 || i == y {
					str += "-"
					// z01.PrintRune('')
				} else if j == 0 || j == x {
					str += "|"
					// z01.PrintRune('')
				} else {
					str += " "
					//z01.PrintRune(32)
				}
			}
			str += "\n"
			//z01.PrintRune(10)
		}
	}
	return str
}

func Raid1bStr(x, y int) string {
	str := ""
	x--
	y--
	if x >= 0 && y >= 0 {
		if y == 0 {
			for j := 0; j <= x; j++ {
				if j == 0 {
					str += "/"
					// z01.PrintRune('')
				} else if j == x {
					str += "\\"
					// z01.PrintRune('')
				} else {
					str += "*"
					//z01.PrintRune('')
				}
			}
			str += "\n"
			//z01.PrintRune(10)
		} else if x == 0 {
			for i := 0; i <= y; i++ {
				if i == 0 {
					str += "/"
					// z01.PrintRune('')
				} else if i == y {
					str += "\\"
					// z01.PrintRune('')
				} else {
					str += "*"
					// z01.PrintRune('')
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		} else {
			for i := 0; i <= y; i++ {
				for j := 0; j <= x; j++ {
					if i == y && j == 0 ||
						i == 0 && j == x {
						str += "\\"
						// z01.PrintRune('')
					} else if i == 0 && j == 0 ||
						i == y && j == x {
						str += "/"
						// z01.PrintRune('')
					} else if i == 0 || i == y {
						str += "*"
						// z01.PrintRune('')
					} else if j == 0 || j == x {
						str += "*"
						//z01.PrintRune('')
					} else {
						str += " "
						// z01.PrintRune(32)
					}
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		}
	}
	return str
}

func Raid1cStr(x, y int) string {
	str := ""
	x--
	y--
	// s1 := 'A'
	// s2 := 'C'
	// s3 := 'B'
	if x >= 0 && y >= 0 {
		if y == 0 {
			for j := 0; j <= x; j++ {
				if j == 0 || j == x {
					str += "A"
					// z01.PrintRune(s1)
				} else {
					str += "B"
					// z01.PrintRune(s3)
				}
			}
			str += "\n"
			//z01.PrintRune(10)
		} else if x == 0 {
			for i := 0; i <= y; i++ {
				if i == 0 {
					str += "A"
					// z01.PrintRune(s1)
				} else if i == y {
					str += "C"
					// z01.PrintRune(s2)
				} else {
					str += "B"
					// z01.PrintRune(s3)
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		} else {
			for i := 0; i <= y; i++ {
				for j := 0; j <= x; j++ {
					if i == y && j == 0 ||
						i == y && j == x {
						str += "C"
						// z01.PrintRune(s2)
					} else if i == 0 && j == 0 ||
						i == 0 && j == x {
						str += "A"
						// z01.PrintRune(s1)
					} else if i == 0 || i == y {
						str += "B"
						// z01.PrintRune(s3)
					} else if j == 0 || j == x {
						str += "B"
						// z01.PrintRune(s3)
					} else {
						str += " "
						// z01.PrintRune(32)
					}
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		}
	}
	return str
}

func Raid1dStr(x, y int) string {
	str := ""
	x--
	y--
	//s1 := 'A'
	//s2 := 'C'
	//s3 := 'B'
	if x >= 0 && y >= 0 {
		if y == 0 {
			for j := 0; j <= x; j++ {
				if j == 0 {
					str += "A"
					//z01.PrintRune(s1)
				} else if j == x {
					str += "C"
					//z01.PrintRune(s2)
				} else {
					str += "B"
					//z01.PrintRune(s3)
				}
			}
			str += "\n"
			//z01.PrintRune(10)
		} else if x == 0 {
			for i := 0; i <= y; i++ {
				if i == 0 || i == y {
					str += "A"
					//z01.PrintRune(s1)
				} else {
					str += "B"
					//z01.PrintRune(s3)
				}
				str += "\n"
				//z01.PrintRue(10)
			}
		} else {
			for i := 0; i <= y; i++ {
				for j := 0; j <= x; j++ {
					if i == 0 && j == x ||
						i == y && j == x {
						str += "C"
						//z01.PrintRune(s2)
					} else if i == 0 && j == 0 ||
						i == y && j == 0 {
						str += "A"
						//z01.PrintRune(s1)
					} else if i == 0 || i == y {
						str += "B"
						//z01.PrintRune(s3)
					} else if j == 0 || j == x {
						str += "B"
						//z01.PrintRune(s3)
					} else {
						str += " "
						//z01.PrintRune(32)
					}
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		}
	}
	return str
}

func Raid1eStr(x, y int) string {
	str := ""
	x--
	y--
	//s1 := 'A'
	//s2 := 'C'
	//s3 := 'B'
	if x >= 0 && y >= 0 {
		if y == 0 {
			for j := 0; j <= x; j++ {
				if j == 0 {
					str += "A"
					//z01.PrintRune(s1)
				} else if j == x {
					str += "C"
					//z01.PrintRune(s2)
				} else {
					str += "B"
					//z01.PrintRune(s3)
				}
			}
			str += "\n"
			//z01.PrintRune(10)
		} else if x == 0 {
			for i := 0; i <= y; i++ {
				if i == 0 {
					str += "A"
					//z01.PrintRune(s1)
				} else if i == y {
					str += "C"
					//z01.PrintRune(s2)
				} else {
					str += "B"
					//z01.PrintRune(s3)
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		} else {
			for i := 0; i <= y; i++ {
				for j := 0; j <= x; j++ {
					if i == 0 && j == x ||
						i == y && j == 0 {
						str += "C"
						//z01.PrintRune(s2)
					} else if i == 0 && j == 0 ||
						i == y && j == x {
						str += "A"
						//z01.PrintRune(s1)
					} else if i == 0 || i == y {
						str += "B"
						//z01.PrintRune(s3)
					} else if j == 0 || j == x {
						str += "B"
						//z01.PrintRune(s3)
					} else {
						str += " "
						//z01.PrintRune(32)
					}
				}
				str += "\n"
				//z01.PrintRune(10)
			}
		}
	}
	return str
}
