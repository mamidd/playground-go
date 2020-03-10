package main

import "fmt"

func main() {
	//modo 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	//modo 2
	// var colors map[string]string

	//modo 3
	// colors := make(map[string]string)

	//cambio valore
	colors["red"] = "black"

	//cancello un valore
	delete(colors, "green")

	fmt.Printf("%v", colors)
}
