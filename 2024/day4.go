package twentyfour

import (
	"bufio"
	"os"
	"strings"
)

const word_size = 4

func XMASCount(input [][]string) int {
	count := 0
	for i, inner := range input {
		for j := range inner {
			if input[i][j] == "X" {
				count += circular_check(i, j, input)
			}
		}
	}
	return count
}

func circular_check(i, j int, input [][]string) int {
	count := 0
	dirs := [8][2]int{
		{-1, 0}, {-1, 1},
		{0, 1}, {1, 1},
		{1, 0}, {1, -1},
		{0, -1}, {-1, -1},
	}

	for _, dir := range dirs {
		word := input[i][j]
		x := i
		y := j
		for range 3 {
			x += dir[1]
			y += dir[0]
			if x < 0 || x >= len(input[i]) || y < 0 || y >= len(input) {
				break
			}
			word += input[x][y]
		}
		if word == "XMAS" {
			count++
		}
	}
	return count
}

func Day4Input() [][]string {
	f, err := os.Open("2024/input4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	result := [][]string{}
	for s.Scan() {
		letters := strings.Split(s.Text(), "")
		result = append(result, letters)
	}
	return result
}
