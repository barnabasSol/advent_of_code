package main

import (
	"fmt"

	twentyfour "github.com/barnabasSol/advent/2024"
)

func main() {
	input, _ := twentyfour.Day3Input()
	// input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	fmt.Printf(
		"twentyfour.CleanCorruptedSum(input): %v\n",
		twentyfour.CleanCorruptedWithConditions(input),
	)

}
