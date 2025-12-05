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
	for scaner.Scan() {
		line := scaner.Text()
		var id int
		fmt.Sscanf(line, "%d", &id)
		for _, r := range freshRanges {
			if id >= r[0] && id <= r[1] {
				countFresh++
				break
			}
		}
	}
	fmt.Println(countFresh)
}
