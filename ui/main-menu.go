package ui

import (
	"fmt"
)

func DisplayMainMenu() {
	var name string

	fmt.Print("Please make a selection: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
}
