// read input file input.txt and find two numbers summing up to 2020, multiply them and print the result
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var checkTokens = [...]string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

//
// read in data and wander through the map
//
func main() {
	passportList := loadData()

	validPassports := len(passportList)

	for _, passport := range passportList {
		for _, elem := range checkTokens {
			if !strings.Contains(passport, elem) {
				fmt.Println(passport)
				validPassports--
				break
			}
		}
	}

	fmt.Printf("amount of valid passports is %d", validPassports)
}

//
// read values from input and return as string slice
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
	var passportData string

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			passportData = passportData + " " + strings.TrimSpace(line)
		} else {
			result = append(result, passportData)
			passportData = ""
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
