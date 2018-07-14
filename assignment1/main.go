package main

import "fmt"

func main() {
	tri := triangle{
		base:   5.0,
		height: 10.0,
	}
	sq := square{
		length: 7.0,
	}

	fmt.Println(printArea(tri))
	fmt.Println(printArea(sq))
}
