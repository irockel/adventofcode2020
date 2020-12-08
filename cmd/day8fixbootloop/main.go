// read input file input.txt and find the boot loop inside the code
// and sum them up.
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
// read in data and run the program to find the boot lopp
//
func main() {
	program := readProgram()
	accState := -1
	changedOps := make([]bool, len(program))

	for accState == -1 {
		accState = findOpBeforeLoop(program, changedOps)
	}

	fmt.Printf("acc right before endless looping: %d\n", accState)
}

//
// read values from input and return as string slices so to have a proper program listing
//
func readProgram() []string {
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
	program := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		program = append(program, strings.TrimSpace(line))

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return program
}

//
// execute the program
//
func findOpBeforeLoop(program []string, changedOps []bool) int {
	var accState, currentOp int
	var flipped bool

	execState := make([]bool, len(program))

	for {
		if execState[currentOp] {
			accState = -1
			break
		}
		execState[currentOp] = true

		if len(program[currentOp]) > 0 {
			stmt := strings.Split(program[currentOp], " ")

			switch stmt[0] {
			case "jmp":
				if !changedOps[currentOp] && !flipped {
					changedOps[currentOp] = true
					flipped = true
					currentOp++
				} else {
					val, _ := strconv.Atoi(stmt[1])
					currentOp += val
				}

			case "nop":
				if !changedOps[currentOp] && !flipped {
					changedOps[currentOp] = true
					flipped = true
					val, _ := strconv.Atoi(stmt[1])
					currentOp += val
				} else {
					currentOp++
				}
			case "acc":
				val, _ := strconv.Atoi(stmt[1])
				accState += val
				currentOp++
			}
		} else {
			fmt.Println("Terminated normally!")
			break
		}
	}

	return accState

}
