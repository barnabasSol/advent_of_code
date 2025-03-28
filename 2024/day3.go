package twentyfour

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func CleanCorruptedWithConditions(input string) int {
	bulk := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\([0-9]{1,3},[0-9]{1,3}\)`).FindAllString(input, -1)
	sum := 0
	do := true
	for _, item := range bulk {
		if strings.HasPrefix(item, "don't()") {
			do = false
			continue
		} else if strings.HasPrefix(item, "do()") {
			do = true
			continue
		}
		if do {
			if strings.HasPrefix(item, "mul") {
				close := strings.IndexAny(item, ")")
				num_str := strings.Split(item[4:close], ",")
				sum += (to_int(num_str[0]) * to_int(num_str[1]))
			}
		}
	}
	return sum
}

func CleanCorruptedSum(input string) int {
	valid_instructions := regexp.
		MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`).
		FindAllString(input, -1)

	sum := 0
	for _, instr := range valid_instructions {
		close := strings.IndexAny(instr, ")")
		num_str := strings.Split(instr[4:close], ",")
		sum += (to_int(num_str[0]) * to_int(num_str[1]))
	}
	return sum
}

func Day3Input() (string, error) {
	f, err := os.Open("2024/input3.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	symbs := ""
	for s.Scan() {
		symbs += s.Text()
	}
	return symbs, nil
}
