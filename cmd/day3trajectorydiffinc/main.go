// read input file input.txt and find two numbers summing up to 2020, multiply them and print the result
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//
// read in data and wander through the map
//
func main() {
	flightMap := loadData()

	var countedTrees int64
	countedTrees = int64(countTrees(flightMap, 1, 1))
	countedTrees = countedTrees * int64(countTrees(flightMap, 3, 1))
	countedTrees = countedTrees * int64(countTrees(flightMap, 5, 1))
	countedTrees = countedTrees * int64(countTrees(flightMap, 7, 1))
	countedTrees = countedTrees * int64(countTrees(flightMap, 1, 2))

	fmt.Printf("the trees counted are %d\n", countedTrees)
}

func countTrees(flightMap []string, xInc int, yInc int) int {
	var countedTrees, x int
	lineLength := len(strings.TrimSpace(flightMap[0]))

	for y := yInc; y < len(flightMap); y += yInc {
		x += xInc

		if x >= lineLength {
			x -= lineLength
		}

		if flightMap[y][x] == '#' {
			countedTrees++
		}
	}

	return countedTrees

}

//
// read values from input and return as string slice
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
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(line) > 0 {
			result = append(result, line)
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
