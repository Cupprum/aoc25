package main

import (
	"fmt"
	"os"
	"strings"
)

func iterateTwo(mp []string, x, y int, cache [][]int) int {
	if y == len(mp)-1 {
		return 1
	}

	switch mp[y][x] {
	case '^':
		cc := 0
		if x-1 >= 0 {
			cc += iterateTwo(mp, x-1, y, cache)
		}
		if x+1 < len(mp[0]) {
			cc += iterateTwo(mp, x+1, y, cache)
		}
		return cc
	case 'S':
		fallthrough
	case '.':
		mp[y] = mp[y][:x] + "|" + mp[y][x+1:]
		val := iterateTwo(mp, x, y+1, cache)
		cache[y][x] = val

		return val
	case '|':
		return cache[y][x]
	}

	return -1
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")
	cache := make([][]int, len(mp))
	for i := range cache {
		cache[i] = make([]int, len(mp[0]))
	}

	pos := strings.Index(mp[0], "S")
	c := iterateTwo(mp, pos, 0, cache)
	fmt.Println(c)
}
