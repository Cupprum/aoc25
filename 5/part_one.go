package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
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

		ranges = append(ranges, []int{min, max})
	}

	counter := 0

	for _, line := range strings.Split(parts[1], "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				counter++
				break
			}
		}
	}

	fmt.Println(counter)
}
