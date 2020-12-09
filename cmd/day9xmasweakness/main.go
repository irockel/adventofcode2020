// read input file input.txt and find the first "broken" value
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
// read in data and find the defunct number + the contiguous sum
//
func main() {
	cryptedData := readData()

	firstDefunct := findDefunct(cryptedData)

	fmt.Printf("first defunct value: %d\n", firstDefunct)
}

//
// read values from input and return as int slices
//
func readData() []int {
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
	crypted := []int{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		val, _ := strconv.Atoi(strings.TrimSpace(line))
		crypted = append(crypted, val)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return crypted
}

//
// find the defunct number in the set of input numbers
//
func findDefunct(data []int) int {

	for i := 25; i < len(data); i++ {
		smallerOnes := findNumbersSmaller(data, data[i], i)
		if (!validate(smallerOnes, data[i])) {
			return data[i]
		}
	}
	return -1
}

//
// for validating find the list of numbers smaller the the possible
// defunct candidate
//
func findNumbersSmaller(data []int, number int, index int) []int {
	result := []int{}

	for i := 0; i < index; i++ {
		if (data[i] < number) {
			result = append(result, data[i])
		}
	}

	return result
}

//
// validate the input set of smaller numbers if there's a possible
// tuple summing up to the candidate. If not we have a defunct numer
//
func validate(input []int, number int) bool {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if j != i && input[j] + input[i] == number {
				return true
			}
		}
	}

	return false
}