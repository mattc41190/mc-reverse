package main

import (
	"fmt"
	"os"

	"matthewcale.com/reverse/lib"
)

const (
	test = "test"
)

func main() {
	reversed, err := lib.Reverse(test)
	if err != nil {
		fmt.Println("An error occurred exiting")
		os.Exit(1)
	}
	fmt.Println(reversed)
}
