package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args
	lenArg := 0
	for i := range arg {
		lenArg = i + 1
	}
	if lenArg > 2 {
		fmt.Printf("Too many arguments\n")
	} else {
		if lenArg == 2 {
			fileName := arg[1]
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Printf("open %v: no such file or directory\n", fileName)
			} else {
				fmt.Printf("%s", content)
			}
		} else {
			fmt.Printf("File name missing\n")
		}
	}
}
