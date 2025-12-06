//go:build ignore

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
	count, _ := markAccess(maps)
	fmt.Println(count)
}

var dirs = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

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
				for _, dir := range dirs {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < len(maps) && nj >= 0 && nj < len(maps[i]) && maps[ni][nj] == '@' {
						count++
					}
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
