package main

import (
	"fmt"
	"os"
	"strings"
)

func checkPos(mp []string, x, y int) bool {
	cc := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := x + dx
			ny := y + dy
			if ny < 0 || ny >= len(mp) || nx < 0 || nx >= len(mp[ny]) {
				continue
			}
			if mp[ny][nx] == '@' || mp[ny][nx] == 'x' {
				cc++
			}
		}
	}

	return cc < 4
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	c := 0

	for y, row := range mp {
		for x, char := range row {
			if char == '@' && checkPos(mp, x, y) {
				c++
			}
		}
	}
	fmt.Println(c)
}
