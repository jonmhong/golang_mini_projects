package main

import "fmt"

func main() {
	// option 1
	// var colors map[string]string

	// option 2
	// colors := make(map[string]string)

	// option 3
	// colors := map[string]string{}

	// declaring a map, where all keys are type string, values type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"
	delete(colors, "white")
	colors["black"] = "#000000"

	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
