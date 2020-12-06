package main

import (
	"fmt"
	"os"

	"github.com/kalensk/advent/days"
)

func main() {
	fmt.Println("Starting advent of code")
	fmt.Println()

	err := days.Day2()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
