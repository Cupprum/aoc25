package main

import (
	"fmt"
	"os"
	"strings"
)

func iterate(mp []string, x, y int) {
	if y == len(mp)-1 {
		mp[y] = mp[y][:x] + "|" + mp[y][x+1:]
		return
	}

	if mp[y][x] == '^' {
		if x-1 >= 0 && mp[y][x-1] != '|' {
			iterate(mp, x-1, y)
		}
		if x+1 < len(mp[0]) && mp[y][x+1] != '|' {
			iterate(mp, x+1, y)
		}
		return
	} else {
		mp[y] = mp[y][:x] + "|" + mp[y][x+1:]
		iterate(mp, x, y+1)
		return
	}
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")
	pos := strings.Index(mp[0], "S")
	iterate(mp, pos, 0)

	c := 0
	for y, v := range mp {
		for x, r := range v {
			if r == '^' {
				if y-1 >= 0 && mp[y-1][x] == '|' {
					c++
				}
			}
		}
	}

	fmt.Println(c)
}
