package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var g1, g2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			g1 = append(g1, num1)
			g2 = append(g2, num2)
		}
	}
	println(similarity_score(g1, g2))
}
