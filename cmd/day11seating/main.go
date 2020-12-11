// read input file input.txt and try to stabilize seat occupation
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//
// read in data, occupy the seats until it stabilizes and count the occupied seats.
//
func main() {
	seatMap := readData()

	finalMap := occupySeats(seatMap)

	fmt.Printf("Occupied seats in the end: %d\n", countAllOccupied(finalMap))
}

//
// occupy the seats in several rounds until no more changes happen.
//
// The following rules are applied to every seat simultaneously:
//
// - If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// - If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
// - Otherwise, the seat's state does not change.
// - Floor (.) never changes; seats don't move, and nobody sits on the floor.
//
func occupySeats(seatMap []string) []string {
	newSeatMap := seatMap
	currentSeatMap := seatMap
	modifiedSeats := 1

	for modifiedSeats != 0 {
		currentSeatMap = newSeatMap

		newSeatMap = []string{}
		modifiedSeats = 0
		for posY, line := range currentSeatMap {
			newLine := ""
			for posX, token := range line {
				if token == 'L' && countOccupied(currentSeatMap, posX, posY) == 0 {
					newLine = newLine + "#"
					modifiedSeats++
				} else if token == '#' && countOccupied(currentSeatMap, posX, posY) >= 4 {
					newLine = newLine + "L"
					modifiedSeats++
				} else {
					newLine = newLine + string(token)
				}
			}
			newSeatMap = append(newSeatMap, newLine)
		}
	}

	return newSeatMap
}

// 
// cound the adjacents seats if they are occupied
// 
func countOccupied(seatMap []string, posX int, posY int) int {
	var occupied int
	// check above current line
	if posY > 0 {
		start := posX - 1
		if start < 0 {
			start = 0
		}
		end := posX + 2
		if end > len(seatMap[posY-1]) {
			end = len(seatMap[posY-1])
		}

		for _, elem := range seatMap[posY-1][start:end] {
			if elem == '#' {
				occupied++
			}
		}
	}

	if posX > 0 && seatMap[posY][posX-1] == '#' {
		occupied++
	}

	if posX < len(seatMap[posY])-1 && seatMap[posY][posX+1] == '#' {
		occupied++
	}

	// check below current line
	if posY < len(seatMap)-1 {
		start := posX - 1
		if start < 0 {
			start = 0
		}
		end := posX + 2
		if end > len(seatMap[posY+1]) {
			end = len(seatMap[posY+1])
		}

		for _, elem := range seatMap[posY+1][start:end] {
			if elem == '#' {
				occupied++
			}
		}
	}

	return occupied
}

//
// count all occupied seats
//
func countAllOccupied(seatMap []string) int {
	var count int

	for _, line := range seatMap {
		for _, token := range line {
			if token == '#' {
				count++
			}
		}
	}

	return count
}

//
// read values from input and return as int slices
//
func readData() []string {
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
	seatMap := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			seatMap = append(seatMap, strings.TrimSpace(line))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return seatMap
}
