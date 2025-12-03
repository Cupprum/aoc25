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
	lines := strings.Split(string(f), "\n")

	counter := 0

	for _, line := range lines {
		j := ""

		l := line
		var r rune

		for i := 11; i >= 0; i-- {
			r, l = findHighest(l, i)
			j += string(r)
		}
		n, err := strconv.Atoi(j)
		if err != nil {
			panic(err)
		}

		counter += n
	}
	fmt.Println(counter)
}
