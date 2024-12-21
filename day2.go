package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func safe_count(records [][]int) int {
	count := 0
	for _, record := range records {
		is_safe := true
		is_increasing := false
		is_decreasing := false
		for j := 0; j < len(record); j++ {
			if j+1 == len(record) {
				break
			}
			if !is_valid_gap(record[j], record[j+1]) {

				is_safe = false
				break
			}
			if record[j] == record[j+1] {
				is_safe = false
				break
			}
			if record[j] > record[j+1] {
				is_decreasing = true
			} else {
				is_increasing = true
			}
			if is_increasing && is_decreasing {
				is_safe = false
				break
			}
		}
		if is_safe {
			count++
		}
	}
	return count
}

func is_valid_gap(x, y int) bool {
	if abs(x-y) > 3 {
		return false
	}
	return true

}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func getInput(file *os.File, row [][]int) [][]int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		level := []int{}
		for _, v := range numbers {
			num, _ := strconv.Atoi(v)
			level = append(level, num)
		}
		row = append(row, level)
	}
	return row
}
