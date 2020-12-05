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
	boardingPasses := loadData()
	var highestSeatID int

	for _, boardingPass := range boardingPasses {
		seatID := calculateSeatID(boardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Printf("the highest seat number is %d\n", highestSeatID)
}

//
// calculate the column, row and seat ID from the boardingPass
//
func calculateSeatID(boardingPass string) int {
	rowRange := 128
	rowRangeLower := 0
	colRange := 8
	colRangeLower := 0

	for _, token := range boardingPass {
		if token == 'F' {
			rowRange /= 2
		} else if token == 'B' {
			rowRange /= 2
			rowRangeLower += rowRange
		} else if token == 'L' {
			colRange /= 2
		} else if token == 'R' {
			colRange /= 2
			colRangeLower += colRange
		}
	}
	seatID := rowRangeLower*8 + colRangeLower

	//fmt.Printf("boardingPass %s has col %d, row %d and seatID %d\n", boardingPass, rowRangeLower, colRangeLower, seatID)
	return seatID
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
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			result = append(result, strings.TrimSpace(line))
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
