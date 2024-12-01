package main

import (
	"fmt"
	"math"

	"slices"
	"strconv"
	"strings"

	"github.com/glokta1/advent-of-code-2024/internal"
)

func main() {
	lines := internal.ReadLines("input")
	var left, right []int
	freq := make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		freq[second]++

		left = append(left, first)
		right = append(right, second)
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum, totalSimilarityScore int
	for i := 0; i < len(left); i++ {
		totalSimilarityScore += left[i] * freq[left[i]]
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(sum, totalSimilarityScore)
}
