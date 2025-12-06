//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MathOperand [4]int

func main() {
	var operands []MathOperand
	var operators []func(a MathOperand) int
	scaner := bufio.NewScanner(os.Stdin)
	rowIndex := 0
	for scaner.Scan() {
		line := scaner.Text()
		rows := strings.Fields(line)
		if rowIndex < 4 {
			if operands == nil {
				operands = make([]MathOperand, len(rows))
			}
			for i, row := range rows {
				n, err := strconv.Atoi(row)
				if err != nil {
					log.Fatal(err)
				}
				operands[i][rowIndex] = n
			}
		} else {
			if operators == nil {
				operators = make([]func(a MathOperand) int, len(rows))
			}
			for i, row := range rows {
				switch row {
				case "+":
					operators[i] = func(a MathOperand) int {
						return a[0] + a[1] + a[2] + a[3]
					}
				case "*":
					operators[i] = func(a MathOperand) int {
						return a[0] * a[1] * a[2] * a[3]
					}
				}
			}
		}
		rowIndex++
	}
	var result int
	for i, operand := range operands {
		result += operators[i](operand)
	}
	fmt.Println(result)
}
