package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scaner.Scan() {
		digits := []rune(scaner.Text())
		leftMax := 0
		leftMaxIndex := -1
		for i := 0; i < len(digits)-1; i++ {
			if int(digits[i]-'0') > leftMax {
				leftMax = int(digits[i] - '0')
				leftMaxIndex = i
			}
		}
		rightMax := 0
		for i := leftMaxIndex + 1; i < len(digits); i++ {
			if int(digits[i]-'0') > rightMax {
				rightMax = int(digits[i] - '0')
			}
		}
		joltage := leftMax*10 + rightMax
		sum += joltage
	}
	fmt.Println(sum)
}
