package main

import (
	"fmt"
	"slices"
)

func preorderTwo(m map[string][]string, curr string, finish string, limits []string, memo map[string]int) int {
	if val, ok := memo[curr]; ok {
		return val
	}

	if curr == finish {
		return 1
	} else if slices.Contains(limits, curr) {
		return 0
	}

	counter := 0
	for _, v := range m[curr] {
		counter += preorderTwo(m, v, finish, limits, memo)
	}

	memo[curr] = counter
	return counter
}

func partTwo() {
	m := generateMap()

	svrToDac := preorderTwo(m, "svr", "dac", []string{"fft", "out"}, map[string]int{})
	fmt.Println("done svrToDac")
	dacToFft := preorderTwo(m, "dac", "fft", []string{"out"}, map[string]int{})
	fmt.Println("done dacToFft")
	fftToOut := preorderTwo(m, "fft", "out", []string{"dac"}, map[string]int{})
	fmt.Println("done fftToOut")

	svrToFft := preorderTwo(m, "svr", "fft", []string{"dac", "out"}, map[string]int{})
	fmt.Println("done svrToFft")
	fftToDac := preorderTwo(m, "fft", "dac", []string{"out"}, map[string]int{})
	fmt.Println("done fftToDac")
	dacToOut := preorderTwo(m, "dac", "out", []string{"fft"}, map[string]int{})
	fmt.Println("done dacToOut")

	fmt.Println(svrToDac, dacToFft, fftToOut)
	fmt.Println(svrToFft, fftToDac, dacToOut)
	fmt.Println(svrToDac*dacToFft*fftToOut + svrToFft*fftToDac*dacToOut)
}
