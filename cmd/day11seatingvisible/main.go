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
// - If a seat is empty (L) and there are no occupied seats which can be seen in each of the eight directions to it,
//   the seat becomes occupied.
// - If a seat is occupied (#) and five or more seats seen from it are also occupied, the seat becomes empty.
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
				} else if token == '#' && countOccupied(currentSeatMap, posX, posY) >= 5 {
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

	// x-1, y-1
	occupied += traverse(seatMap, posX, posY, -1, -1)

	// x, y-1
	occupied += traverse(seatMap, posX, posY, 0, -1)

	// x+1, y-1
	occupied += traverse(seatMap, posX, posY, +1, -1)

	// x+1, y
	occupied += traverse(seatMap, posX, posY, +1, 0)

	// x+1, y+1
	occupied += traverse(seatMap, posX, posY, +1, +1)

	// x, y+1
	occupied += traverse(seatMap, posX, posY, 0, +1)

	// x-1, y+1
	occupied += traverse(seatMap, posX, posY, -1, +1)

	// x, y-1
	result := traverse(seatMap, posX, posY, -1, 0)
	occupied += result

	return occupied
}

//
// traverse the seat map in the given inc direction and check if there's
// a occupied seat. A empty seat blocks the view to the occupied seat
// so the traverse ends.
//
func traverse(seatMap []string, posX, posY, incX, incY int) int {
	travX := posX + incX
	travY := posY + incY

	for (travX >= 0 && travX < len(seatMap[0])) && (travY >= 0 && travY < len(seatMap)) {
		if seatMap[travY][travX] == 'L' {
			return 0
		} else if seatMap[travY][travX] == '#' {
			return 1
		}

		travX += incX
		travY += incY
	}

	return 0
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
