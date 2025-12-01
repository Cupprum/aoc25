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

	pos := 50
	counter := 0

	for _, line := range lines {
		letter := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		counter += number / 100

		for i := 0; i < number%100; i++ {
			switch letter {
			case 'L':
				pos--
				if pos < 0 {
					pos = 99
				}
			case 'R':
				pos++
				if pos > 99 {
					pos = 0
				}
			}

			if pos == 0 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
