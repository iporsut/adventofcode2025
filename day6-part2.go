//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OperatorInfo struct {
	start  int
	end    int
	opRune rune
}

type MathOperand []string

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	var inputs []string
	for scaner.Scan() {
		line := scaner.Text()
		inputs = append(inputs, line)
	}
	operatorsLine := []rune(inputs[len(inputs)-1])
	operandsLines := inputs[:len(inputs)-1]
	var operatorInfos []OperatorInfo
	var i int
	for i < len(operatorsLine) {
		var start int
		var end int
		if operatorsLine[i] != ' ' {
			start = i
			i++
			for i < len(operatorsLine) && operatorsLine[i] == ' ' {
				i++
			}
			end = i
			opRune := operatorsLine[start]
			operatorInfos = append(operatorInfos, OperatorInfo{start, end, opRune})
		}
	}
	operands := make([]MathOperand, len(operatorInfos))
	for i, opInfo := range operatorInfos {
		operands[i] = make(MathOperand, len(operandsLines))
		for j, line := range operandsLines {
			operands[i][j] = line[opInfo.start:opInfo.end]
		}
	}
	var result int
	for i, op := range operands {
		oper := operatorInfos[i]
		var maxLen int
		for _, val := range op {
			if len(val) > maxLen {
				maxLen = len(val)
			}
		}
		var opInts []int
		for j := maxLen - 1; j >= 0; j-- {
			var digits []rune
			for _, val := range op {
				if j < len(val) {
					digits = append(digits, rune(val[j]))
				}
			}
			if len(strings.TrimSpace(string(digits))) == 0 {
				continue
			}
			var num int
			fmt.Sscanf(string(digits), "%d", &num)
			opInts = append(opInts, num)
		}
		switch oper.opRune {
		case '+':
			var sum int
			for _, n := range opInts {
				sum += n
			}
			result += sum
		case '*':
			prod := 1
			for _, n := range opInts {
				prod *= n
			}
			result += prod
		}
	}
	fmt.Println(result)
}
