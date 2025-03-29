package twentyfour

import (
	"bufio"
	"os"
	"strings"
)

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
