package loops

import "fmt"

func Map() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	for color, hex := range colors {
		fmt.Printf("Color: %s, Hex: %s\n", color, hex)
	}
	fmt.Println("Map Loop completed.")
}
