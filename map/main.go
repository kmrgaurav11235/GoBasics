package main

import "fmt"

func main() {
	// Creating Map: Method 1
	colors1 := map[string]string{
		"red":   " #FF0000",
		"green": "#00FF00", // note the comma in this line too
	}
	fmt.Println(colors1)

	// Creating Map: Method 2
	colors2 := make(map[string]string)

	// Adding to map
	colors2["yellow"] = "#FFFF00"
	colors2["orange"] = "#FFA500"
	colors2["purple"] = "#800080"

	fmt.Println(colors2)

	// Deleting from map
	delete(colors1, "green")
	fmt.Println(colors1)

	printMap(colors2)
}

func printMap(c map[string]string) {
	// Iterating over maps
	for color, hex := range c {
		fmt.Println("Hexcode for color", color, "is", hex)
	}
}
