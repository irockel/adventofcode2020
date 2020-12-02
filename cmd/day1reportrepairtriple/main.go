// read input file input.txt and find two numbers summing up to 2020, multiply them and print the result
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
// read in data and search and calculate result
//
func main() {
	values, _ := readData()

	for k := range values {
		diff := 2020 - k

		val1, val2 := calculate(diff, values)

		if val1 != -1 {
			fmt.Printf("values are %d, %d, %d\n", k, val1, val2)
			fmt.Printf("result is %d\n", (k * val1 * val2))
			break
		}
	}
}

//
// try to find a fitting tuple for the given boundary
//
func calculate(boundary int, values map[int]bool) (val1 int, val2 int) {
	val1, val2 = -1, -1

	for k := range values {
		diff := boundary - k

		if values[diff] {
			val1 = k
			val2 = diff
			break
		}
	}

	return val1, val2
}

//
// read values from input and return as key map
//
func readData() (values map[int]bool, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	result := make(map[int]bool)

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		val, error := strconv.Atoi(strings.TrimSpace(line))
		if error == nil {
			result[val] = true
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}
	return result, nil
}
