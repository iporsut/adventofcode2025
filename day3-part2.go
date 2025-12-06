//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	sum := new(big.Int)
	for scaner.Scan() {
		digits := []rune(scaner.Text())
		var joltageDigits [12]rune
		var nextIndex int

		for i := 12; i > 0; i-- {
			if len(digits) == (i - 1) {
				copy(joltageDigits[nextIndex:], digits)
				break
			}
			leftPart := digits[:len(digits)-(i-1)]
			leftMax := 0
			leftMaxIndex := -1
			for j := len(leftPart) - 1; j >= 0; j-- {
				if int(leftPart[j]-'0') >= leftMax {
					leftMax = int(leftPart[j] - '0')
					leftMaxIndex = j
				}
			}
			joltageDigits[nextIndex] = rune(leftMax + '0')
			nextIndex++
			digits = digits[leftMaxIndex+1:]
		}
		joltage := new(big.Int)
		joltage, _ = joltage.SetString(string(joltageDigits[:]), 10)
		sum = sum.Add(sum, joltage)
	}
	fmt.Println(sum)
}
