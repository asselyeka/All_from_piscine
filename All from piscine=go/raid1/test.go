package main

import (
	"fmt"
	"os"
	"strconv"

	student ".."
)

func main() {
	arg := os.Args

	if len(arg) == 3 {
		x, err1 := strconv.Atoi(arg[1])
		y, err2 := strconv.Atoi(arg[2])
		if err1 != nil || err2 != nil {
			fmt.Println(err1.Error())
			fmt.Println(err2.Error())
			return
		}
		student.Raid1b(x, y)
	}
}
