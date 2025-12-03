package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findHighest(s string, ign int) (rune, string) {
	m := rune(s[0])
	pos := 0
	for p, c := range s {
		if p > len(s)-1-ign {
			break
		}
		if c > m {
			m = c
			pos = p
		}
	}
	return m, s[pos+1:]
}

func partOne() {
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

		for i := 1; i >= 0; i-- {
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
