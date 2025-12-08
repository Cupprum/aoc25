package main

import (
	"fmt"
	"os"
	"strings"
)

func partTwo() {
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

	for len(dists) != 0 {
		i := findLowest(dists)

		groups = addToGroups(paths[i], groups)

		if len(groups) == 1 && len(groups[0]) == len(nums) {
			nums := strings.Split(paths[i], "-")
			x1, _, _ := parseNum(nums[0])
			x2, _, _ := parseNum(nums[1])
			fmt.Println(x1 * x2)
			return
		}

		dists = append(dists[:i], dists[i+1:]...)
		paths = append(paths[:i], paths[i+1:]...)
	}
}
