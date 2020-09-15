package main

import "fmt"

func main() {
	// straight forward syntax
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#008000",
	}
	// fmt.Println(colors)
	printMap(colors)

	//declare and initialize later
	colors2 := make(map[string]string)
	colors2["white"] = "#FFFFFF"
	colors2["green"] = "#008000"
	delete(colors2, "green")
	fmt.Println(colors2)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " - " + hex)
	}
}
