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

//
// run the specified program
//
func runProgram(programCode []string) int {
	registers := make(map[int]int)
	memPattern := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
	memSum := 0
	var mask string

	for _, codeLine := range programCode {
		if strings.HasPrefix(codeLine, "mem") {
			groups := memPattern.FindAllStringSubmatch(codeLine, -1)[0]
			register, _ := strconv.Atoi(groups[1])
			value, _ := strconv.Atoi(groups[2])

			changedRegisters := applyMask(mask, register)

			for _, register := range changedRegisters {
				registers[register] = value
			}
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
// apply the (meta) mask and calculate all registers
//
func applyMask(mask string, value int) []int {
	upperBound := 0
	conv := strconv.FormatInt(int64(value), 2)
	convLen := len(conv)
	registers := []int{}

	for i := 0; i < 36-convLen; i++ {
		conv = "0" + conv
	}

	masked := ""
	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			masked = masked + "X"
		} else if mask[i] == '1' {
			masked = masked + "1"
		} else if mask[i] == '0' {
			masked = masked + string(conv[i])
		}
	}

	for _, token := range masked {
		if token == 'X' {
			if upperBound == 0 {
				upperBound = 2
			} else {
				upperBound *= 2
			}

		}
	}

	for i := 0; i < upperBound; i++ {
		bits := strconv.FormatInt(int64(i), 2)
		register := ""
		bitsIndex := len(bits) - 1
		for j := len(masked) - 1; j >= 0; j-- {
			val := masked[j]
			if val == 'X' {
				if bitsIndex < 0 {
					register = "0" + register
				} else {
					register = string(bits[bitsIndex]) + register
				}
				bitsIndex--
			} else {
				register = string(val) + register
			}
		}
		registerVal, _ := strconv.ParseInt(register, 2, 64)
		registers = append(registers, int(registerVal))
	}

	return registers
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
