package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var maps [][]rune
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		line := scaner.Text()
		maps = append(maps, []rune(line))
	}
	count := markAccess(maps)
	fmt.Println(count)
}

func markAccess(maps [][]rune) int {
	countAll := 0

	for i := 0; i < len(maps); i++ {
		for j := 0; j < len(maps[i]); j++ {
			if maps[i][j] == '.' {
				continue
			}

			if maps[i][j] == '@' {
				count := 0
				if i-1 >= 0 && j-1 >= 0 && maps[i-1][j-1] == '@' {
					count++
				}
				if i-1 >= 0 && maps[i-1][j] == '@' {
					count++
				}
				if i-1 >= 0 && j+1 < len(maps[i]) && maps[i-1][j+1] == '@' {
					count++
				}
				if j+1 < len(maps[i]) && maps[i][j+1] == '@' {
					count++
				}
				if i+1 < len(maps) && j+1 < len(maps[i]) && maps[i+1][j+1] == '@' {
					count++
				}
				if i+1 < len(maps) && maps[i+1][j] == '@' {
					count++
				}
				if i+1 < len(maps) && j-1 >= 0 && maps[i+1][j-1] == '@' {
					count++
				}
				if j-1 >= 0 && maps[i][j-1] == '@' {
					count++
				}

				if count < 4 {
					countAll++
				}
			}
		}
	}

	return countAll
}
