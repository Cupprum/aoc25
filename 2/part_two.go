package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rngs := strings.Split(string(f), ",")

	counter := 0

	for _, rng := range rngs {
		nums := strings.Split(rng, "-")
		lowS, highS := nums[0], nums[1]

		lN, err := strconv.Atoi(lowS)
		if err != nil {
			panic(err)
		}
		hN, err := strconv.Atoi(highS)
		if err != nil {
			panic(err)
		}

		for i := lN; i <= hN; i++ {
			if isInvalidID(i) {
				counter += i
			}
		}
	}

	fmt.Println(counter)
}

func isInvalidID(n int) bool {
	s := strconv.Itoa(n)

	// Potential lenghts
	for l := 1; l <= len(s)/2; l++ {

		// Skip if the total length isn't divisible by the pattern length
		if len(s)%l != 0 {
			continue
		}

		if strings.Repeat(s[:l], len(s)/l) == s {
			return true
		}
	}

	return false
}
