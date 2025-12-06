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

	s := []int{}

	ops := strings.Fields(lines[len(lines)-1])

	for _, o := range ops {
		switch o {
		case "+":
			s = append(s, 0)
		case "*":
			s = append(s, 1)
		}
	}

	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}

		vals := strings.Fields(line)
		for j, v := range vals {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			switch ops[j] {
			case "+":
				s[j] += n
			case "*":
				s[j] *= n
			}
		}
	}

	c := 0
	for _, v := range s {
		c += v
	}

	fmt.Println(c)

}
