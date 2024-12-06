package main

import "slices"

func distance_sum(g1, g2 []int) int {
	distances := make([]int, len(g1))
	slices.Sort(g1)
	slices.Sort(g2)
	distances[0] = g1[0] - g2[0]
	for i := 0; i < len(g1); i++ {
		distances[i] = g1[i] - g2[i]
		if distances[i] < 0 {
			distances[i] *= -1
		}
	}
	sum := 0
	for _, v := range distances {
		sum += v
	}
	return sum
}

func similarity_score(g1, g2 []int) int {
	var scores = make([]int, len(g1))
	for i, l := range g1 {
		score := 0
		for _, r := range g2 {
			if l == r {
				score++
			}
		}
		scores[i] = l * score
	}
	sum := 0
	for _, v := range scores {
		sum += v
	}
	return sum
}
