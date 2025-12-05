package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findOverlap(ranges [][]int, min, max int) [][]int {
	overlap := [][]int{}

	for _, r := range ranges {
		if (min >= r[0] && min <= r[1]) || (max >= r[0] && max <= r[1]) || (min <= r[0] && max >= r[1]) {
			overlap = append(overlap, r)
		}
	}

	return overlap
}

func findIndex(ranges [][]int, target []int) int {
	for i, r := range ranges {
		if r[0] == target[0] && r[1] == target[1] {
			return i
		}
	}
	return -1
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(f), "\n\n")

	ranges := [][]int{}

	for _, rng := range strings.Split(parts[0], "\n") {
		lines := strings.Split(rng, "-")

		min, err := strconv.Atoi(lines[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(lines[1])
		if err != nil {
			panic(err)
		}

		overlap := findOverlap(ranges, min, max)
		for _, o := range overlap {
			i := findIndex(ranges, o)
			ranges = append(ranges[:i], ranges[i+1:]...)
		}
		if len(overlap) > 0 {
			newMin := min
			newMax := max
			for _, o := range overlap {
				if o[0] < newMin {
					newMin = o[0]
				}
				if o[1] > newMax {
					newMax = o[1]
				}
			}
			ranges = append(ranges, []int{newMin, newMax})
		}

		if len(overlap) == 0 {
			ranges = append(ranges, []int{min, max})
		}
	}

	counter := 0
	for _, r := range ranges {
		counter += r[1] - r[0] + 1
	}
	fmt.Println(counter)
}
