// read input file input.txt and find the 3's and 1's in the jolts
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

//
// read in data, sort the list of ints and count number of threes and ones
//
func main() {
	jolts := readData()

	sort.Ints(jolts)

	threes, ones := countThreesAndOnes(jolts)

	// the adapter is per definition +3 to the highest jolt
	threes++

	fmt.Printf("Product of 1s's and 3's: %d\n", threes*ones)
}

//
// count threes and ones
//
func countThreesAndOnes(jolts []int) (int, int) {
	var threes, ones int

	for i := 0; i < len(jolts)-1; i++ {
		if jolts[i+1]-jolts[i] == 1 {
			ones++
		} else {
			threes++
		}
	}

	return threes, ones
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
	jolts := []int{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		val, _ := strconv.Atoi(strings.TrimSpace(line))
		jolts = append(jolts, val)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return jolts
}
