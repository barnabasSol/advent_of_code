package main

import (
	"fmt"

	twentyfour "github.com/barnabasSol/advent/2024"
)

func main() {
	input := twentyfour.Day4Input()
	// input := [][]string{
	// 	{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
	// 	{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
	// 	{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
	// 	{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
	// 	{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
	// 	{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
	// 	{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
	// 	{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
	// 	{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
	// 	{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	// }

	fmt.Printf("twentyfour.CrossMASCount(input): %v\n", twentyfour.CrossMASCount(input))

}
