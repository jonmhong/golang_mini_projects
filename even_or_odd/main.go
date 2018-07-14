package main

import (
	"fmt"
)

func main() {
	intRange := intSlice(10)

	for _, i := range intRange {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
