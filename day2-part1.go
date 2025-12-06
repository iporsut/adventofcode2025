//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		line := scaner.Text()
		pairs := strings.Split(line, ",")
		var sum int
		for _, r := range pairs {
			bounds := strings.Split(r, "-")
			var start, end int
			fmt.Sscanf(bounds[0], "%d", &start)
			fmt.Sscanf(bounds[1], "%d", &end)

			for i := start; i <= end; i++ {
				s := fmt.Sprint(i)
				if len(s)%2 == 1 {
					continue
				}
				mid := len(s) / 2
				firstHalf := s[:mid]
				secondHalf := s[mid:]
				if firstHalf == secondHalf {
					sum += i
				}
			}
		}
		fmt.Println(sum)
	}
}
