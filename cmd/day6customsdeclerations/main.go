// read input file input.txt and calculate all boarding passes to find the one
// with the highest ID.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//
// read in data and wander through the slices and check for the highest
// seat ID
//
func main() {
	groupCustomsForms := loadData()

	sumCustomsForms := calculateSums(groupCustomsForms)

	fmt.Printf("custom forms sum up to %d\n", sumCustomsForms)
}

//
// calculate sum of customs forms
func calculateSums(sumCustomsForms []string) int {
	var sum int
	for _, groupCustoms := range sumCustomsForms {
		sum += len(groupCustoms)
	}

	return sum
}

//
// read values from input and return as string slices
//
func loadData() []string {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil
	}
	defer file.Close()

	var result []string

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var groupCustoms string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) == 0 {
			result = append(result, strings.TrimSpace(groupCustoms))
			groupCustoms = ""
		} else {
			for _, resp := range strings.TrimSpace(line) {
				if !strings.Contains(groupCustoms, string(resp)) {
					groupCustoms = groupCustoms + string(resp)
				}
			}
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return result
}
