// read input file input.txt and calculate the results.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//
// read in data, calculate the results and print the sum
//
func main() {
	calculations := readCalculations()

	results := calculateResults(calculations)

	sum := sumResult(results)

	fmt.Printf("The sum of results: %d\n", sum)
}

//
// read values from input and return calculations as string slices
//
func readCalculations() []string {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	calculations := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			calculations = append(calculations, strings.TrimSpace(line))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return calculations
}

//
// calculate the results of the given calculations and return
// as list
//
func calculateResults(calculations []string) []int {
	results := []int{}

	for _, calc := range calculations {
		_, calc := calculate(calc, 0)
		results = append(results, calc)
	}

	return results
}

//
// do the actual math and return the result of each calculation
// the operation is recursive
//
func calculate(calc string, pos int) (int, int) {
	result := 0
	operand := 0 // 0 addition, 1 multiplication
	digits := ""
	stack := []int{} // stack needed for operator precedence

	for pos < len(calc) {
		switch calc[pos] {
		case '(':
			subResult := 0
			pos, subResult = calculate(calc, pos+1)
			switch operand {
			case 0:
				result += subResult
			case 1:
				result *= subResult
			}
		case '+':
			operand = 0
			pos++
		case '*':
			operand = 1
			stack = append(stack, result)
			result = 1
			pos++
		case ')':
			if len(digits) > 0 {
				result = flush(digits, operand, result)
			}
			for _, val := range stack {
				result *= val
			}
			return pos + 1, result
		case ' ':
			if len(digits) > 0 {
				result = flush(digits, operand, result)
			}

			digits = ""
			pos++
		default:
			digits = digits + string(calc[pos])
			pos++
		}

	}
	if len(digits) > 0 {
		result = flush(digits, operand, result)
	}

	for _, val := range stack {
		result *= val
	}

	return pos, result
}

//
// flush the current digits and add to the result
//
func flush(digits string, operand int, result int) int {
	//fmt.Printf("flush digits: %s, result: %d, operand: %d\n", digits, result, operand)
	var val int
	val, _ = strconv.Atoi(digits)
	switch operand {
	case 0:
		return result + val
	case 1:
		return result * val
	}

	return 0
}

//
// sum up the given results
//
func sumResult(results []int) int {
	result := 0
	for _, val := range results {
		result += val
	}

	return result
}
