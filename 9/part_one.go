package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}

type Point struct {
	x int
	y int
}

func calculate(A, B Point) int {
	x := A.x - B.x
	y := A.y - B.y

	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}

	x += 1
	y += 1

	return x * y
}

func partOne() {
	f := must(os.ReadFile("input.txt"))
	pointsS := strings.Split(string(f), "\n")

	points := []Point{}

	for _, point := range pointsS {
		splitted := strings.Split(point, ",")
		points = append(points, Point{
			x: must(strconv.Atoi(splitted[0])),
			y: must(strconv.Atoi(splitted[1])),
		})
	}

	max := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			size := calculate(points[i], points[j])
			if size > max {
				max = size
			}
		}
	}
	fmt.Println(max)
}
