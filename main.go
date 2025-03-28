package main

import (
	"log"
	"time"

	twentyfour "github.com/barnabasSol/advent/2024"
)

func main() {
	start := time.Now()
	// reports := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// }
	reports := twentyfour.Day2Input()

	log.Println(twentyfour.SafeCountWithDampner(reports))
	log.Println(time.Since(start))
}
