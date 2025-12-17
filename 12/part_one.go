package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

func partOne() {
	f := must(os.ReadFile("input.txt"))

	input := strings.Split(string(f), "\n\n")
	lines := input[len(input)-1]

	counter := 0

	for _, line := range strings.Split(lines, "\n") {
		pts := strings.Split(line, ": ")

		sizes := strings.Split(pts[0], "x")
		width := must(strconv.Atoi(sizes[0]))
		length := must(strconv.Atoi(sizes[1]))

		canFit := (width / 3) * (length / 3)

		packages := 0
		for _, v := range strings.Split(pts[1], " ") {
			packages += must(strconv.Atoi(v))
		}

		if packages <= canFit {
			counter++
		}
	}

	fmt.Println(counter)
}
