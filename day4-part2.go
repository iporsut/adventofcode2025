package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var maps [][]rune
	var countAll, count int
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		line := scaner.Text()
		maps = append(maps, []rune(line))
	}

	for {
		count, maps = markAccess(maps)
		countAll += count
		if count == 0 {
			break
		}
	}

	fmt.Println(countAll)
}

func markAccess(maps [][]rune) (int, [][]rune) {
	result := make([][]rune, len(maps))
	countAll := 0

	for i := 0; i < len(maps); i++ {
		result[i] = make([]rune, len(maps[i]))
		for j := 0; j < len(maps[i]); j++ {
			if maps[i][j] == '.' {
				result[i][j] = '.'
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
					result[i][j] = '.'
					countAll++
				} else {
					result[i][j] = '@'
				}
			}
		}
	}

	return countAll, result
}
