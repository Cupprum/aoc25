package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseNum(i string) (int, int, int) {
	vals := strings.Split(i, ",")

	x := must(strconv.Atoi(vals[0]))
	y := must(strconv.Atoi(vals[1]))
	z := must(strconv.Atoi(vals[2]))

	return x, y, z
}

func distance(x1, y1, z1, x2, y2, z2 int) int {
	v := math.Pow(float64(x1-x2), 2) +
		math.Pow(float64(y1-y2), 2) +
		math.Pow(float64(z1-z2), 2)

	return int(math.Sqrt(v))
}

func findLowest(distances []int) int {
	index := 0
	lowest := distances[0]
	for i, v := range distances {
		if v < lowest {
			lowest = v
			index = i
		}
	}
	return index
}

func addToGroups(numsS string, groups [][]string) [][]string {
	nums := strings.Split(numsS, "-")

	aI, bI := -1, -1
	a := nums[0]
	b := nums[1]

	for i, group := range groups {
		if slices.Contains(group, a) {
			aI = i
		}
		if slices.Contains(group, b) {
			bI = i
		}
	}

	if aI == -1 && bI == -1 {
		// Not present
		groups = append(groups, []string{a, b})
	} else if aI == -1 && bI != -1 {
		// First is not present
		groups[bI] = append(groups[bI], a)
	} else if aI != -1 && bI == -1 {
		// Second is not present
		groups[aI] = append(groups[aI], b)
	} else if aI != bI {
		// Both are not present, merging
		groups[aI] = append(groups[aI], groups[bI]...)
		groups = append(groups[:bI], groups[bI+1:]...)
	}

	return groups
}

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}

func partOne() {
	f := must(os.ReadFile("input.txt"))
	nums := strings.Split(string(f), "\n")

	dists := make([]int, 0)
	paths := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		x1, y1, z1 := parseNum(nums[i])
		for j := i + 1; j < len(nums); j++ {
			x2, y2, z2 := parseNum(nums[j])
			dists = append(dists, distance(x1, y1, z1, x2, y2, z2))
			paths = append(paths, nums[i]+"-"+nums[j])
		}
	}

	groups := [][]string{}

	// Change to 1000 on larger input
	for j := 0; j < 10; j++ {
		i := findLowest(dists)

		groups = addToGroups(paths[i], groups)

		dists = append(dists[:i], dists[i+1:]...)
		paths = append(paths[:i], paths[i+1:]...)
	}

	lens := []int{}
	for _, group := range groups {
		lens = append(lens, len(group))
	}
	slices.Sort(lens)

	fmt.Println(lens[len(lens)-1] * lens[len(lens)-2] * lens[len(lens)-3])
}
