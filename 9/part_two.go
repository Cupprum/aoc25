package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partTwo() {
	// Load input
	f := must(os.ReadFile("input.txt"))
	pointsS := strings.Split(string(f), "\n")

	points := []Point{}

	for _, point := range pointsS {
		splitted := strings.Split(point, ",")

		x := must(strconv.Atoi(splitted[0]))
		y := must(strconv.Atoi(splitted[1]))

		points = append(points, Point{x: x, y: y})
	}

	// Calculate areas
	areas := map[int][]Point{}
	indexes := []int{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := calculate(points[i], points[j])

			indexes = append(indexes, area)
			areas[area] = []Point{points[i], points[j]}
		}
	}

	// Sort areas
	slices.SortFunc(indexes, func(a, b int) int { return b - a })

	// Generate vectors
	vectors := [][]Point{}
	for i := 0; i < len(points)-1; i++ {
		if i == 0 {
			vectors = append(vectors, []Point{points[len(points)-1], points[0]})
		} else {
			vectors = append(vectors, []Point{points[i], points[i+1]})
		}
	}

	// Iterate through areas and check for collisions with vectors
	for _, area := range indexes {
		rec := areas[area]

		xMin, xMax, yMin, yMax := 0, 0, 0, 0

		if rec[0].x < rec[1].x {
			xMin = rec[0].x
			xMax = rec[1].x
		} else {
			xMin = rec[1].x
			xMax = rec[0].x
		}

		if rec[0].y < rec[1].y {
			yMin = rec[0].y
			yMax = rec[1].y
		} else {
			yMin = rec[1].y
			yMax = rec[0].y
		}

		colides := false
		for _, v := range vectors {
			if v[0].x == v[1].x && xMin < v[0].x && v[0].x < xMax {
				if yMin < v[0].y && v[0].y < yMax && yMin < v[1].y && v[1].y < yMax {
					// Both are inside - This is probably fine
				} else if (yMin < v[0].y && v[0].y < yMax) || (yMin < v[1].y && v[1].y < yMax) {
					// One inside
					colides = true
					break
				} else if (v[0].y <= yMin && v[1].y <= yMin) || (yMax <= v[0].y && yMax <= v[1].y) {
					// Both above or under - This is probably fine
				} else {
					// Both outside
					colides = true
					break
				}
			} else if v[0].y == v[1].y && yMin < v[0].y && v[0].y < yMax {
				if xMin < v[0].x && v[0].x < xMax && xMin < v[1].x && v[1].x < xMax {
					// Both are inside - This is probably fine
				} else if (xMin < v[0].x && v[0].x < xMax) || (xMin < v[1].x && v[1].x < xMax) {
					// One inside
					colides = true
					break
				} else if (v[0].x <= xMin && v[1].x <= xMin) || (xMax <= v[0].x && xMax <= v[1].x) {
					// Both above or under - This is probably fine
				} else {
					// Both outside
					colides = true
					break
				}
			}
		}
		if !colides {
			fmt.Println(area)
			return
		}
	}
}
