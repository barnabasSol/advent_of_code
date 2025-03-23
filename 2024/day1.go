package twentyfour

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Distance(a, b []int) int {
	slices.Sort(a)
	slices.Sort(b)
	sum := 0
	for i := range a {
		val := a[i] - b[i]
		if val < 0 {
			val *= -1
		}
		sum += val
	}
	return sum
}

func Similarity(left, right []int) int {
	sum := 0
	for _, l := range left {
		count := 0
		for _, r := range right {
			if l == r {
				count += 1
			}
		}
		sum += l * count
	}
	return sum
}

func Day1Input() ([]int, []int) {
	f, err := os.Open("2024/input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		seps := strings.Split(line, "   ")
		left = append(left, to_int(seps[0]))
		right = append(right, to_int(seps[1]))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return left, right
}

func to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
