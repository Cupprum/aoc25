package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func splitNum(s string) string {
	midpoint := len(s) / 2
	return s[:midpoint]
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rngs := strings.Split(string(f), ",")

	counter := 0

	for _, rng := range rngs {
		nums := strings.Split(rng, "-")
		lowS, highS := nums[0], nums[1]

		// IDs should not start with leading zeros
		lowS = strings.TrimLeft(lowS, "0")
		highS = strings.TrimLeft(highS, "0")

		lN, err := strconv.Atoi(lowS)
		if err != nil {
			panic(err)
		}
		hN, err := strconv.Atoi(highS)
		if err != nil {
			panic(err)
		}

		lsN, err := strconv.Atoi(splitNum(lowS))
		if err != nil {
			lsN, err = strconv.Atoi(lowS)
			if err != nil {
				panic(err)
			}
		}
		hsN, err := strconv.Atoi(splitNum(highS))
		if err != nil {
			panic(err)
		}

		if len(lowS)%2 != 0 {
			lsN = int(math.Pow(10, float64(len(lowS)/2)))
		}
		if len(highS)%2 != 0 {
			hsN *= 10
		}

		for i := lsN; i <= hsN; i++ {
			iS, err := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(i))
			if err != nil {
				panic(err)
			}

			if iS > hN {
				break
			}

			if iS >= lN && iS <= hN {
				counter += iS
			}

		}
	}

	fmt.Println(counter)

}
