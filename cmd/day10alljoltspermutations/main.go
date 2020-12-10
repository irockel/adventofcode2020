// read input file input.txt and determine all possible permutations
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

//
// read in data, sort the list and determine all possible permutations
//
func main() {
	jolts := readData()

	sort.Ints(jolts)

	tuples := findAllTuples(jolts)

	combinations := calculateCombinations(tuples)

	fmt.Printf("All possible combinations: %d\n", combinations)
}

//
// find all tuples allowing multiple combinations
//
func findAllTuples(jolts []int) [][]int {
	tuples := [][]int{}
	tuple := []int{jolts[0]}

	for i := 1; i < len(jolts); i++ {
		if jolts[i-1] == jolts[i]-1 {
			tuple = append(tuple, jolts[i])
		} else {
			if len(tuple) > 1 {
				tuples = append(tuples, tuple)
			}
			tuple = []int{}
			tuple = append(tuple, jolts[i])
		}
	}

	if len(tuple) > 1 {
		tuples = append(tuples, tuple)
	}

	return tuples
}

// 
// calculate the possible combinations using the binomial coefficient
//
func calculateCombinations(tuples [][]int) int64 {
	var result int64
	ommitOne := new(big.Int)
	ommitTwo := new(big.Int)

	result = 1

	for i := 0; i < len(tuples); i++ {
		tupleLen := len(tuples[i])
		if tupleLen > 3 {

			ommitOne.Binomial(int64(tupleLen-2), int64(tupleLen-3))
			ommitTwo.Binomial(int64(tupleLen-2), int64(tupleLen-4))
			result *= (ommitOne.Int64() + ommitTwo.Int64() + 1)
		} else if tupleLen == 3 {
			result *= 2
		}
	}

	return result
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
	jolts := []int{0}

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
