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
// describes a bus line
type busInfo struct {
	busID     int
	increment int
	delta     int
}

//
// read in data, check the bus table for the earliest possible depart time and
// multiply bus id with waiting time
//
func main() {
	busTables := readBusTables()

	fmt.Println(busTables)

	startPoint := 0
	for bound := 2; bound <= len(busTables); bound++ {
		highestPos, highestBusID := findHighestBusID(busTables[:bound])
		if startPoint == 0 {
			startPoint = highestBusID
		}
		calculateDeltas(busTables[:bound], highestPos)
		fmt.Println(busTables)
		fmt.Println(startPoint)
		startPoint = earlistCascade(busTables[:bound], startPoint) - highestBusID
	}

	fmt.Printf("earlist start point all busses depart one after another: %d\n", startPoint)
}

//
// find the earliest point all buses depart from another properly
//
func earlistCascade(busTables []busInfo, increment int) int {
	// no sense in starting below 100000000000000 (hint from puzzle)
	time := 0

	found := false
	for time = 0; !found; time += increment {
		found = true
		for _, busInfo := range busTables {

			if busInfo.delta < 0 {
				if time%busInfo.busID != busInfo.delta*-1 {
					found = false
					break
				}
			} else if busInfo.delta > 0 {
				if time%busInfo.busID != busInfo.busID-busInfo.delta {
					found = false
					break
				}
			}

		}
	}

	return time
}

func calculateDeltas(busTables []busInfo, highestPos int) {
	// first all lower the highest bus ID
	delta := 0
	for i := highestPos; i >= 0; i-- {
		busTables[i].delta = delta
		delta -= busTables[i].increment
	}

	// now all higher than the highest bus ID
	delta = 0
	for i := highestPos + 1; i < len(busTables); i++ {
		delta += busTables[i].increment
		busTables[i].delta = delta
	}
}

func findHighestBusID(busTables []busInfo) (int, int) {
	var highestPos, highestBusID int

	for pos, busInfo := range busTables {
		if busInfo.busID > highestBusID {
			highestPos = pos
			highestBusID = busInfo.busID
		}
	}

	return highestPos, highestBusID
}

//
// read values from input and return as int slices
//
func readBusTables() []busInfo {
	fmt.Println("reading inputtest.txt")

	file, err := os.Open("./inputtest.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	busTables := []busInfo{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 && strings.Contains(line, ",") {
			busTablesS := strings.Split(line, ",")
			inc := 0
			for _, busID := range busTablesS {
				inc++
				if busID != "x" {
					busInt, _ := strconv.Atoi(strings.TrimSpace(busID))
					if len(busTables) == 0 {
						// the first step does not count
						inc--
					}
					busTables = append(busTables, busInfo{busInt, inc, 0})
					inc = 0
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

	return busTables
}
