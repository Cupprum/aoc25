package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setCounter(ops []string, iter int) int {
	switch ops[iter] {
	case "+":
		return 0
	case "*":
		return 1
	}
	return -1
}

func getNum(lines []string, col int) int {
	num := ""
	for r := 0; r < len(lines)-1; r++ {
		if lines[r][col] != ' ' {
			num += string(lines[r][col])
		}
	}
	if num != "" {
		val, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		return val
	}
	return -1
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")

	ops := strings.Fields(lines[len(lines)-1])
	iter := len(ops) - 1

	total := 0
	counter := setCounter(ops, iter)

	for i := len(lines[0]) - 1; i >= 0; i-- {
		val := getNum(lines, i)

		if val == -1 {
			iter--
			total += counter
			counter = setCounter(ops, iter)
			continue
		}

		switch ops[iter] {
		case "+":
			counter += val
		case "*":
			counter *= val
		}
	}

	total += counter
	fmt.Println(total)
}
