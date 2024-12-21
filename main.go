package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	rows := [][]int{
		// {7, 6, 4, 2, 1},
		// {1, 2, 7, 8, 9},
		// {9, 7, 6, 2, 1},
		// {1, 3, 2, 4, 5},
		// {8, 6, 4, 4, 1},
		// {1, 3, 6, 7, 9},
	}
	input := getInput(file, rows)

	count := safe_count(input)
	println(count)
}
