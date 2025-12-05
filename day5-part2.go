package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range [2]int

func main() {
	var freshRanges []Range
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		line := scaner.Text()
		if line == "" {
			break
		}

		var r Range
		fmt.Sscanf(line, "%d-%d", &r[0], &r[1])
		freshRanges = append(freshRanges, r)
	}
	var countFresh int
	var mergedFreshRanges []Range
	for i := 0; i < len(freshRanges); i++ {
		merged := false
		for j := i + 1; j < len(freshRanges); j++ {
			if freshRanges[j][0] > freshRanges[i][1] || freshRanges[j][1] < freshRanges[i][0] {
				continue
			}
			freshRanges[j] = Range{
				min(freshRanges[i][0], freshRanges[j][0]),
				max(freshRanges[i][1], freshRanges[j][1]),
			}
			merged = true
		}
		if !merged {
			mergedFreshRanges = append(mergedFreshRanges, freshRanges[i])
		}
	}
	for _, r := range mergedFreshRanges {
		countFresh += r[1] - r[0] + 1
	}
	fmt.Println(countFresh)
}
