package main

import (
	"fmt"
	"os"
)

func main() {
	str := []string{"01", "galaxy", "galaxy 01"}
	arguments := []string(os.Args)
	lenA := 0
	for i := range arguments {
		lenA = i + 1 + 1 - 1
	}
	if lenA == 1 {
		return
	}
	arg := []string(arguments[1:])
	for _, argI := range arg {
		for _, strI := range str {
			if strI == argI {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
