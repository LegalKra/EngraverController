package main

import (
	"fmt"
	"tejnecky/engravercontroller/movement"
	"tejnecky/engravercontroller/usb"
)

func main() {
	// Usb comunication
	// TODO FIX windows compile error
	// C/MSYS2 compiler problem ?
	usb.ListAwaiablePorts()
	// Test of movement
	origin, gCode := movement.HomePosition()
	fmt.Println(origin, "\n", gCode)
	origin, gCode = movement.StartMovementFullSpeed(origin, 5.3, 2.7)
	fmt.Println(origin, "\n", gCode)
	origin, gCode = movement.ContinueMovmentX(origin, -5.0)
	fmt.Println(origin, "\n", gCode)

}
