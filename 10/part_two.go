package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var memo = make(map[string]int)

func parseJoltage(s string) []int {
	s = strings.Trim(s, "{}")
	vals := strings.Split(s, ",")

	r := make([]int, len(vals))
	for i, v := range vals {
		r[i] = must(strconv.Atoi(v))
	}
	return r
}

// parseCoefficients reads the button schematics and converts them into a coefficient matrix.
// The output is a matrix where coeffs[i][j] is 1 if button i affects counter j, 0 otherwise.
func parseCoefficients(buttonStrings []string, numCounters int) [][]int {
	coeffs := make([][]int, len(buttonStrings))

	for i, s := range buttonStrings {
		s = strings.Trim(s, "()")

		// Map of affected indices to quickly check for presence
		affected := make(map[int]bool)

		// Handle the case where the button string is not empty
		if s != "" {
			indices := strings.Split(s, ",")
			for _, indexStr := range indices {
				index := must(strconv.Atoi(indexStr))
				if index < numCounters {
					affected[index] = true
				}
			}
		}

		// Build the coefficient vector for this button
		coeffs[i] = make([]int, numCounters)
		for j := 0; j < numCounters; j++ {
			if affected[j] {
				coeffs[i][j] = 1
			}
		}
	}
	return coeffs
}

func generatePatterns(coeffs [][]int) map[string]int {
	numButtons := len(coeffs)
	if numButtons == 0 {
		return nil
	}
	numCounters := len(coeffs[0])

	// Key: string representation of the resulting sum vector (e.g., "1,0,1,1"). Value: min presses (0 or 1).
	patternCosts := make(map[string]int)

	// Iterate through all 2^N combinations using a bitmask
	for i := 0; i < 1<<numButtons; i++ {
		currentPattern := make([]int, numCounters)
		presses := 0

		for j := 0; j < numButtons; j++ {
			if (i>>j)&1 == 1 { // If j-th button is pressed
				presses++
				for k := 0; k < numCounters; k++ {
					currentPattern[k] += coeffs[j][k]
				}
			}
		}

		var sb strings.Builder
		for k, v := range currentPattern {
			if k > 0 {
				sb.WriteRune(',')
			}
			sb.WriteString(strconv.Itoa(v))
		}
		patternKey := sb.String()

		if cost, ok := patternCosts[patternKey]; !ok || presses < cost {
			patternCosts[patternKey] = presses
		}
	}

	return patternCosts
}

func solveSingleAux(coeffs [][]int, goal []int, patternCosts map[string]int) int {
	allZero := true
	for _, g := range goal {
		if g != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	var sb strings.Builder
	for i, g := range goal {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(g))
	}
	goalKey := sb.String()

	if result, ok := memo[goalKey]; ok {
		return result
	}

	const impossible = 1_000_000
	minPresses := impossible

	for patternKey, patternCost := range patternCosts {
		patternStrVals := strings.Split(patternKey, ",")
		pattern := make([]int, len(patternStrVals))
		for i, s := range patternStrVals {
			pattern[i] = must(strconv.Atoi(s))
		}

		isValid := true
		newGoal := make([]int, len(goal))
		for i := 0; i < len(goal); i++ {
			p, g := pattern[i], goal[i]

			// Pattern must not exceed goal and parity must match (g-p must be even)
			if p > g || (g-p)%2 != 0 {
				isValid = false
				break
			}
			newGoal[i] = (g - p) / 2
		}

		if isValid {
			recursiveResult := solveSingleAux(coeffs, newGoal, patternCosts)

			// Only update if the subproblem is solvable
			if recursiveResult < impossible {
				currentTotal := patternCost + 2*recursiveResult
				if currentTotal < minPresses {
					minPresses = currentTotal
				}
			}
		}
	}

	memo[goalKey] = minPresses
	return minPresses
}

func partTwo() {
	f := must(os.ReadFile("input.txt"))
	lines := strings.Split(string(f), "\n")

	counter := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		joltageGoal := parseJoltage(parts[len(parts)-1])
		buttonStrings := parts[1 : len(parts)-1]
		coeffs := parseCoefficients(buttonStrings, len(joltageGoal))

		memo = make(map[string]int)
		patternCosts := generatePatterns(coeffs)

		minPresses := solveSingleAux(coeffs, joltageGoal, patternCosts)

		counter += minPresses
	}

	fmt.Println(counter)
}
