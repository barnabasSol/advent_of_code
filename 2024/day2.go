package twentyfour

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func damp(input []int, i, j int) []int {
	out := []int{}
	copy(out, input)

}

func IsSafeWithDampner(reports [][]int) int {
	safe := 0
	for _, x := range reports {
		leng := len(x)
		damp_count := 0
		slices.Sort(x)
		for j := range leng - 1 {
			if j == leng {
				break
			}
			check := x[j] - x[j+1]
			if check*-1 > 3 || check == 0 {
				damp_count += 1
				damp(x, j, j+1)
			}

		}
	}
	return safe
}

func IsSafe(reports [][]int) int {
	safe := 0
	for _, x := range reports {
		leng := len(x)
		is_increasing, is_decreasing := false, false
		for j := range leng - 1 {
			if j == leng {
				break
			}
			check := x[j] - x[j+1]
			if check < 0 {
				is_increasing = true
			} else if check > 0 {
				is_decreasing = true
			} else {
				safe -= 1
				break
			}
			if is_increasing && is_decreasing {
				safe -= 1
				break
			}
			if is_increasing {
				if check*-1 > 3 {
					safe -= 1
					break
				}
			} else if is_decreasing {
				if check > 3 {
					safe -= 1
					break
				}
			}
		}
		safe += 1
	}
	return safe
}

func Day2Input() [][]int {
	f, err := os.Open("2024/input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	reports := [][]int{}
	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		rep_int := []int{}
		for _, rep_str := range report {
			rep_int = append(rep_int, to_int(rep_str))
		}
		reports = append(reports, rep_int)

	}
	return reports
}
