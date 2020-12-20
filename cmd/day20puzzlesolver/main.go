// read input file input.txt and find all messages matching rule "0"
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
// read in data, calculate the results and print the sum
//
func main() {
	tiles := readImageTiles()

	image := buildImage(tiles)

	edgeValues := calculateEdgeValues(image)

	fmt.Printf("The amout of messages complying with rule 0 are: %d\n", len(results))
}

type tileDef struct {
	data []string
	left struct {
		val int
		reverse int
	}
	right []int
}

//
// read values from input and image tiles with calculated border
// values as map keys
//
func readImageTiles() (map[int]tileDef) {
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
	tiles := map[int]tileDef{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			if strings.Contains(trimmedLine, "#") || strings.Contains(trimmedLine, "-") {
			} else {
			}
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return tiles
}

