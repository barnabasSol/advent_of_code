package twentyfour

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func SafeCountDampner(nested_reports [][]int) int {
	safe_count := 0
	for _, reports := range nested_reports {
		if is_safe(reports) {
			safe_count += 1
		} else {
			fixed := try_fix(reports)
			if fixed {
				safe_count++
			}
		}
	}
	return safe_count
}

func try_fix(reports []int) bool {
	out := make([]int, len(reports))
	copy(out, reports)
	leng := len(reports)
	for i := range leng {
		out = slices.Delete(out, i, i+1)
		if is_safe(out) {
			return true
		}
		out = make([]int, len(reports))
		copy(out, reports)
	}
	return false
}

func is_safe(reports []int) bool {
	leng := len(reports) - 1
	is_increasing, is_decreasing := false, false
	for i := range leng {
		if i == leng {
			break
		}
		check := reports[i] - reports[i+1]
		if check < 0 {
			is_increasing = true
		} else if check > 0 {
			is_decreasing = true
		} else {
			return false
		}
		if is_increasing && is_decreasing {
			return false
		}
		if is_increasing {
			if check*-1 > 3 {
				return false
			}
		} else if is_decreasing {
			if check > 3 {
				return false
			}
		}
	}
	return true
}

func SafeCount(reports [][]int) int {
	safe := 0
	for _, x := range reports {
		is_increasing, is_decreasing := false, false
		leng := len(x) - 1
		for j := range leng {

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
