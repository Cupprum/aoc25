package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

func compare[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func generateCombinations(t int, d int) [][]int {
	arr := [][]int{}

	if t < d {
		return nil
	}
	cur := make([]int, d)
	for i := 0; i < d; i++ {
		cur[i] = i
	}

	for j := 0; ; j++ {
		// Copy slice, because by default they are passed by reference
		arr = append(arr, slices.Clone(cur))

		i := d - 1
		// Find the rightmost index that can be incremented
		for i >= 0 && cur[i] == t-(d-i) {
			i--
		}

		if i < 0 {
			break
		}

		// Increment and reset the rest of array
		cur[i]++
		for j := i + 1; j < d; j++ {
			cur[j] = cur[j-1] + 1
		}
	}

	return arr
}

func parseTarget(s string) []bool {
	s = strings.Trim(s, "[]")

	r := make([]bool, len(s))
	for i, c := range s {
		switch c {
		case '.':
			r[i] = false
		case '#':
			r[i] = true
		default:
			fmt.Println(string(c))
			panic("Invalid input")
		}
	}

	return r
}

func parseButtons(b []string) [][]int {
	r := [][]int{}

	for i, vals := range b {
		vals = strings.Trim(vals, "()")
		vals := strings.Split(vals, ",")

		r = append(r, []int{})

		for _, v := range vals {
			r[i] = append(r[i], must(strconv.Atoi(string(v))))
		}
	}

	return r
}

func calculate(t []bool, b [][]int, indexes []int) bool {
	current := make([]bool, len(t))

	for i := 0; i < len(t); i++ {
		current[i] = false
	}

	for _, index := range indexes {
		for _, v := range b[index] {
			current[v] = !current[v]
		}
	}

	return compare(t, current)
}

func partOne() {
	f := must(os.ReadFile("input.txt"))
	lines := strings.Split(string(f), "\n")

	counter := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		target := parseTarget(parts[0])
		buttons := parseButtons(parts[1 : len(parts)-1])

	out: // Break out of nested for loop
		for d := 1; d < 10; d++ {
			combinations := generateCombinations(len(buttons), d)

			for _, c := range combinations {
				if calculate(target, buttons, c) {
					counter += d
					break out
				}
			}
		}

	}

	fmt.Println(counter)
}
