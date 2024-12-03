package main

import (
	"flag"
	"fmt"

	"github.com/lgrn/aoc24/day1"
	"github.com/lgrn/aoc24/day2"
)

func main() {
	day := flag.Int("d", 0, "View solution for day N")
	flag.Parse()

	switch *day {
	case 1:
		fmt.Println(day1.Solve())
	case 2:
		fmt.Println(day2.Solve())
	default:
		fmt.Printf("No solution available for day %d.\n", *day)
		fmt.Printf("Use the -d [day] option.\n")
	}
}
