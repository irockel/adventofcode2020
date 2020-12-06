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
func calculateSums(sumCustomsForms []int) int {
	var sum int
	for _, groupCustoms := range sumCustomsForms {
		sum += groupCustoms
	}

	return sum
}

//
// read values from input and return as string slices
//
func loadData() []int {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil
	}
	defer file.Close()

	var result []int

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var memberCount int
	charOffset := 97

	groupCustoms := make([]int, 26)

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) == 0 {
			customsResult := 0
			for i := 0; i < len(groupCustoms); i++ {
				if groupCustoms[i] == memberCount {
					customsResult++
				}
			}
			result = append(result, customsResult)

			// reset data
			memberCount = 0
			groupCustoms = make([]int, 26)
		} else {
			memberCount++
			for _, resp := range strings.TrimSpace(line) {
				groupCustoms[int(resp)-charOffset]++
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
