package main

import (
	"adventofcode2023/day01"
	"adventofcode2023/day02"
	"adventofcode2023/day03"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		var arg = args[0]
		switch arg {
		case "1":
			p1, p2 := day01.Solve()
			fmt.Printf("Solution 1: %s\n", p1)
			fmt.Printf("Solution 2: %s\n", p2)
		case "2":
			p1, p2 := day02.Solve()
			fmt.Printf("Solution 1: %s\n", p1)
			fmt.Printf("Solution 2: %s\n", p2)
		case "3":
			p1, p2 := day03.Solve()
			fmt.Printf("Solution 1: %s\n", p1)
			fmt.Printf("Solution 2: %s\n", p2)
		default:
			fmt.Printf("Value %s is unrecognized.\n", arg)
		}

	} else {
		fmt.Println("Welcome to advent of code 2023!")
	}
}
