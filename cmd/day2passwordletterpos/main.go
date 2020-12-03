// read input file input.txt and find two numbers summing up to 2020, multiply them and print the result
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
// read in data and search and calculate result
//
func main() {
	checkPasswords()
}

//
// stores all information about a given password line with
// the password and the rule
//
type passwordInfo struct {
	minOccur  int
	maxOccur  int
	ruleToken string
	password  string
}

//
// read values from input and return as key map
//
func checkPasswords() {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	var validPasswordCount int

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if strings.Contains(line, ":") {
			val := parseLine(line)

			if ((string(val.password[val.minOccur-1]) == val.ruleToken) && (string(val.password[val.maxOccur-1]) != val.ruleToken)) ||
				((string(val.password[val.minOccur-1]) != val.ruleToken) && (string(val.password[val.maxOccur-1]) == val.ruleToken)) {
				fmt.Println(val)
				validPasswordCount++
			}
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	fmt.Printf("valid passwords: %d\n", validPasswordCount)
}

//
// countRuleLetter: count the letter occurence
func countRuleLetter(val passwordInfo) int {
	var counter int
	for _, char := range val.password {
		if char == rune(val.ruleToken[0]) {
			counter++
		}
	}

	return counter
}

//
// parse the given password line and returned a password info struct
//
func parseLine(line string) passwordInfo {

	lineTokens := strings.Split(line, ":")

	var result passwordInfo
	result.password = strings.TrimSpace(lineTokens[1])

	ruleTokens := strings.Split(lineTokens[0], " ")
	result.ruleToken = ruleTokens[1]

	rangeTokens := strings.Split(ruleTokens[0], "-")

	result.minOccur, _ = strconv.Atoi(rangeTokens[0])
	result.maxOccur, _ = strconv.Atoi(rangeTokens[1])

	return result
}
