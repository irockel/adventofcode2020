// read input file input.txt and find two numbers summing up to 2020, multiply them and print the result
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

var checkTokens = [...]string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

//
// read in data and wander through the map
//
func main() {
	passportList := loadData()

	validPassports := len(passportList)

	for _, passport := range passportList {
		alreadyFalse := false
		for _, elem := range checkTokens {
			if !strings.Contains(passport, elem) {
				validPassports--
				alreadyFalse = true
				break
			}
		}

		if !alreadyFalse {
			passportElems := strings.Split(passport, " ")
			for _, elem := range passportElems {
				if !checkElement(elem) {
					validPassports--
					break
				}
			}
		}
	}

	fmt.Printf("amount of valid passports is %d\n", validPassports)
}

//
// check the validity of the given rules:
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
//
func checkElement(element string) bool {
	elemTokens := strings.Split(element, ":")
	result := true

	switch elemTokens[0] {
	case "byr":
		val, _ := strconv.Atoi(elemTokens[1])
		result = val >= 1920 && val <= 2002
	case "iyr":
		val, _ := strconv.Atoi(elemTokens[1])
		result = val >= 2010 && val <= 2020
	case "eyr":
		val, _ := strconv.Atoi(elemTokens[1])
		result = val >= 2020 && val <= 2030
	case "hgt":
		var validHgt = regexp.MustCompile("(\\d+)(cm|in)")
		result = validHgt.MatchString(elemTokens[1])
		if result {
			var heightTokens = validHgt.FindStringSubmatch(elemTokens[1])
			height, _ := strconv.Atoi(heightTokens[1])
			switch heightTokens[2] {
			case "cm":
				result = height >= 150 && height <= 193
			case "in":
				result = height >= 59 && height <= 76

			}
		}
	case "hcl":
		var validHcl = regexp.MustCompile("#[0-9a-f]{6}")
		result = validHcl.MatchString(elemTokens[1])
	case "ecl":
		var validEcl = regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")
		result = validEcl.MatchString(elemTokens[1])
	case "pid":
		var validPid = regexp.MustCompile("^\\d{9}$")
		result = validPid.MatchString(elemTokens[1])
	}
	return result
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
