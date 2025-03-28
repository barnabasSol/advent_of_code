package twentyfour

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

func damp(reports []int, i, z int) bool {
	out := make([]int, len(reports))
	copy(out, reports)
	out = slices.Delete(out, i, z)
	is_increasing, is_decreasing := false, false
	is_safe := true
	leng := len(out) - 1

	is_safe = damp_check(leng, out, is_increasing, is_decreasing, is_safe)
	if is_safe {
		log.Println("is safe")
		log.Println(reports)
		log.Println(i, z)
		log.Println("//////////////")
		return true
	}

	is_increasing, is_decreasing = false, false
	out = make([]int, len(reports))
	copy(out, reports)
	out = slices.Delete(out, z, z+1)
	is_safe = true
	is_safe = damp_check(leng, out, is_increasing, is_decreasing, is_safe)
	log.Println(is_safe)
	log.Println(reports)
	log.Println(i, z)
	log.Println("//////////////")
	return is_safe
}

func damp_check(
	leng int,
	out []int,
	is_increasing, is_decreasing, is_safe bool,
) bool {
	for j := range leng {
		if j == leng {
			break
		}
		check := out[j] - out[j+1]
		if check < 0 {
			is_increasing = true
		} else if check > 0 {
			is_decreasing = true
		} else {
			is_safe = false
			break
		}
		if is_increasing && is_decreasing {
			is_safe = false
			break
		}
		if is_increasing {
			if check*-1 > 3 {
				is_safe = false
				break
			}
		} else if is_decreasing {
			if check > 3 {
				is_safe = false
				break
			}
		}
	}
	return is_safe
}

func SafeCountWithDampner(reports [][]int) int {
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
				is_safe := damp(x, j, j+1)
				if is_safe {
					break
				}
				safe -= 1
				break
			}
			if is_increasing && is_decreasing {
				is_safe := damp(x, j, j+1)
				if is_safe {
					break
				}
				safe -= 1
				break
			}
			if is_increasing {
				if check*-1 > 3 {
					is_safe := damp(x, j, j+1)
					if is_safe {
						break
					}
					safe -= 1
					break
				}
			} else if is_decreasing {
				if check > 3 {
					is_safe := damp(x, j, j+1)
					if is_safe {
						break
					}
					safe -= 1
					break
				}
			}
		}
		safe += 1
	}
	return safe
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
