// read input file input.txt and navigate the ferry through the storm
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

//
// read in data, navigate through the data and print out the manhattan distance
//
func main() {
	navList := readNavigation()

	x, y := navigate(navList)

	fmt.Printf("Manhattan Distance: %d\n", int(math.Abs(float64(x))+math.Abs(float64(y))))
}

//
// navigate through the list
//
func navigate(navList []string) (int, int) {
	var x, y int
	direction := 90 // east

	for _, elem := range navList {
		action := elem[0]
		val, _ := strconv.Atoi(elem[1:])

		switch action {
		case 'N':
			y += val
		case 'S':
			y -= val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'R':
			direction += val
			if direction >= 360 {
				direction -= 360
			}
		case 'L':
			direction -= val
			if direction < 0 {
				direction += 360
			}
		case 'F':
			switch direction {
			case 0:
				y += val
			case 90:
				x += val
			case 180:
				y -= val
			case 270:
				x -= val
			}
		}
	}

	return x, y
}

//
// read values from input and return as string slices
//
func readNavigation() []string {
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
	navList := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			navList = append(navList, strings.TrimSpace(line))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return navList
}
