package main

import (
	"fmt"
	"os"
	"strings"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func generateMap() map[string][]string {
	f := must(os.ReadFile("input.txt"))
	lines := strings.Split(string(f), "\n")

	m := map[string][]string{}

	for _, line := range lines {
		pts := strings.Split(line, ": ")

		k := pts[0]
		v := strings.Split(pts[1], " ")

		m[k] = v
	}

	return m
}

func preorder(m map[string][]string, curr string) int {
	vals := m[curr]

	if curr == "out" {
		return 1
	}

	counter := 0
	for _, v := range vals {
		counter += preorder(m, v)
	}

	return counter
}

func partOne() {
	m := generateMap()

	fmt.Println(preorder(m, "you"))
}
