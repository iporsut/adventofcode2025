//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	count := 0
	n := 50
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		line := scaner.Text()
		d, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case 'R':
			for i := 0; i < d; i++ {
				n++
				if n > 99 {
					n = 0
				}
				if n == 0 {
					count++
				}
			}
		case 'L':
			for i := 0; i < d; i++ {
				n--
				if n < 0 {
					n = 99
				}
				if n == 0 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
