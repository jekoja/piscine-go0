package main

import "github.com/01-edu/z01"

type Door struct {
	isOpen bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(d *Door) bool {
	PrintStr("Door Opening...")
	d.isOpen = true
	return true
}

func CloseDoor(d *Door) bool {
	PrintStr("Door Closing...")
	d.isOpen = false
	return true
}

func IsDoorOpen(d *Door) bool {
	PrintStr("is the Door opened ?")
	return d.isOpen
}

func IsDoorClose(d *Door) bool {
	PrintStr("is the Door closed ?")
	return !d.isOpen
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.isOpen {
		CloseDoor(door)
	}
}
