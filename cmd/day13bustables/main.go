// read input file input.txt and check for the first possible bus
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
// read in data, check the bus table for the earliest possible depart time and
// multiply bus id with waiting time
//
func main() {
	startPoint, busTables := readBusTables()

	busID, waitTime := checkTables(startPoint, busTables)

	fmt.Printf("Earliest BusID multiplied by the waiting time: %d\n", busID*waitTime)
}

//
// check the startpoint and match it to the given bustables
// return earliest possible bus and waiting time
//
func checkTables(startPoint int, busTables []int) (int, int) {
	var minDiff int
	var bestBusID int
	for _, busID := range busTables {
		diff := busID - startPoint%busID
		if diff < minDiff || minDiff == 0 {
			minDiff = diff
			bestBusID = busID
		}
	}

	return bestBusID, minDiff
}

//
// read values from input and return as int slices
//
func readBusTables() (int, []int) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return -1, nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	busTables := []int{}
	startPoint := 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 && strings.Contains(line, ",") {
			busTablesS := strings.Split(line, ",")
			for _, busID := range busTablesS {
				if busID != "x" {
					busInt, _ := strconv.Atoi(strings.TrimSpace(busID))
					busTables = append(busTables, busInt)
				}
			}
		} else if len(strings.TrimSpace(line)) > 0 {
			startPoint, _ = strconv.Atoi(strings.TrimSpace(line))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return startPoint, busTables
}
