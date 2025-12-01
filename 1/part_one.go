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
	lines := strings.Split(string(f), "\n")

	pos := 50
	counter := 0

	for _, line := range lines {
		letter := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch letter {
		case 'L':
			pos -= number
		case 'R':
			pos += number
		}
		pos %= 100
		if pos < 0 {
			pos += 100
		}

		if pos == 0 {
			counter++
		}
	}

	fmt.Println(counter)
}
