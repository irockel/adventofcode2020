// read input file input.txt and check for the first possible bus
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//
// read in data, check the bus table for the earliest possible depart time and
// multiply bus id with waiting time
//
func main() {
	programCode := readProgram()

	memSum := runProgram(programCode)

	fmt.Printf("Sum of all memory registers: %d\n", memSum)
}

func runProgram(programCode []string) int {
	registers := make(map[int]int)
	memPattern := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
	memSum := 0
	mask := ""

	for _, codeLine := range programCode {
		if strings.HasPrefix(codeLine, "mem") {
			groups := memPattern.FindAllStringSubmatch(codeLine, -1)[0]
			register, _ := strconv.Atoi(groups[1])
			value, _ := strconv.Atoi(groups[2])
			value = applyMask(value, mask)

			registers[register] = value
		} else if strings.HasPrefix(codeLine, "mask") {
			mask = strings.TrimSpace(strings.Split(codeLine, "=")[1])
		}
	}

	for _, register := range registers {
		memSum += register
	}

	return memSum
}

//
// apply the given mask to all the registers. The bits are counted from the
// end til the beginning, so the highest pos in the array is in fact the lowest pos
//
func applyMask(value int, mask string) int {
	for pos, val := range mask {
		bitPos := len(mask) - 1 - pos
		if val == '1' {
			value = setBit(value, uint(bitPos))
		} else if val == '0' {
			value = clearBit(value, uint(bitPos))
		}
	}

	return value
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

//
// read values from input and return as int slices
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
	programCode := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			programCode = append(programCode, line)
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return programCode
}
