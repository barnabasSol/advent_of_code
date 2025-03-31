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
				count += count_valid_dir(i, j, input)

			}
		}
	}
	return count
}

func count_valid_dir(i, j int, input [][]string) int {
	count := 0
	count += count_left(i, j, input)
	count += count_bottom_left(i, j, input)
	count += count_bottom(i, j, input)
	count += count_bottom_right(i, j, input)
	count += count_right(i, j, input)
	count += count_top_right(i, j, input)
	count += count_top(i, j, input)
	count += count_top_left(i, j, input)
	return count
}

func count_left(i, j int, input [][]string) int {
	count := 0
	word := ""
	for range word_size {
		if j < 0 {
			break
		}
		word += input[i][j]
		j -= 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_bottom_left(i, j int, input [][]string) int {
	count := 0
	word := ""
	for range word_size {
		if j < 0 || i >= len(input) {
			break
		}
		word += input[i][j]
		j -= 1
		i += 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_bottom(i, j int, input [][]string) int {
	count := 0
	word := ""
	for range word_size {
		if i >= len(input) {
			break
		}
		word += input[i][j]
		i += 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_bottom_right(i, j int, input [][]string) int {
	count := 0
	word := ""
	i_pos := i
	for range word_size {
		if j >= len(input[i_pos]) || i >= len(input) {
			break
		}
		word += input[i][j]
		i += 1
		j += 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_right(i, j int, input [][]string) int {
	count := 0
	word := ""
	i_pos := i
	for range word_size {
		if j >= len(input[i_pos]) {
			break
		}
		word += input[i][j]
		j += 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_top_right(i, j int, input [][]string) int {
	count := 0
	word := ""
	i_pos := i
	for range word_size {
		if j >= len(input[i_pos]) || i < 0 {
			break
		}
		word += input[i][j]
		j += 1
		i -= 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_top(i, j int, input [][]string) int {
	count := 0
	word := ""
	for range word_size {
		if i < 0 {
			break
		}
		word += input[i][j]
		i -= 1
	}
	if word == "XMAS" {
		count += 1
	}
	return count
}

func count_top_left(i, j int, input [][]string) int {
	count := 0
	word := ""
	for range word_size {
		if i < 0 || j < 0 {
			break
		}
		word += input[i][j]
		i -= 1
		j -= 1
	}
	if word == "XMAS" {
		count += 1
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
