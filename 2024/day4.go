package twentyfour

import (
	"bufio"
	"os"
	"strings"
)

func CrossMASCount(input [][]string) int {
	count := 0
	row_leng := len(input)
	for i := 1; i < row_leng-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] == "A" {
				count += count_mas(input, i, j)
			}
		}
	}
	return count
}

func count_mas(input [][]string, i, j int) int {
	count := 0
	points := [4][2]int{
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}
	word1 := input[i+points[0][1]][j+points[0][0]] + input[i][j] + input[i+points[1][1]][j+points[1][0]]
	word2 := input[i+points[2][1]][j+points[2][0]] + input[i][j] + input[i+points[3][1]][j+points[3][0]]
	if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
		count += 1
	}
	return count
}

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
