package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const OPEN = false
const CLOSE = true

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}

//"Door Opening...\nis the Door closed ?\nis the Door opened ?\nDoor Closing...\n"
