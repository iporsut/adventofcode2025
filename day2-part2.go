//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
				ss := []rune(s)
				mid := len(ss) / 2

				var invalid bool

				for i := 1; i <= mid; i++ {
					var m []rune
					for chunk := range slices.Chunk(ss, i) {
						if m == nil {
							m = chunk
							continue
						}
						if !slices.Equal(m, chunk) {
							m = nil
							break
						}
					}
					if m != nil {
						invalid = true
						break
					}
				}

				if invalid {
					sum += i
				}
			}
		}
		fmt.Println(sum)
	}
}
