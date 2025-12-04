package main

import (
	"fmt"
	"os"
	"strings"
)

func search(mp []string) {
	for y, row := range mp {
		for x, char := range row {
			if char == '@' && checkPos(mp, x, y) {
				mp[y] = mp[y][:x] + "x" + mp[y][x+1:]
			}
		}
	}
}

func cleanup(mp []string) int {
	c := 0

	for y, row := range mp {
		for x, char := range row {
			if char == 'x' {
				mp[y] = mp[y][:x] + "." + mp[y][x+1:]
				c++
			}
		}
	}

	return c
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	c := 0

	for {
		search(mp)
		cc := cleanup(mp)
		if cc == 0 {
			break
		}
		c += cc
	}

	fmt.Println(c)
}
