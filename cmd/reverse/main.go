package main

import (
	"fmt"
	"os"
	"strings"

	"matthewcale.com/reverse/lib"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("reverse requires at least 1 argument")
		os.Exit(1)
	}

	arg := strings.Join(args, " ")

	reversed, err := lib.Reverse(arg)
	if err != nil {
		fmt.Println("An error occurred exiting")
		os.Exit(1)
	}
	fmt.Println(reversed)
}
