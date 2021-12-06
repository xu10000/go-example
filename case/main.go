package main

import "fmt"

const RED int = 1
const GREEN int = 2
const YELLOW int = 3

func main() {
	color := 2
	switch color {
	case RED:
		fmt.Println("------ print", "RED")
	case GREEN, YELLOW:
		fmt.Println("------ print", "GREEN OR YELLOW")
	}

}
