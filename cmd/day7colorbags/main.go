// read input file input.txt and find all possible bag combinations
// and sum them up.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"regexp"
)

//
// read in data and wander through the slices and check for the highest
// seat ID
//
func main() {
	rules := createRules()
	markedColors := make(map[string]bool)

	ruleCount := countRules(rules, markedColors, "shiny gold")

	fmt.Printf("rules with different outer color %d\n", ruleCount)
}

//
// read values from input and return as string slices
//
func createRules() map[string][]string {
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
	rules := make(map[string][]string)

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 && !strings.Contains(line, "no other bags") {
			ruleColors := extractColors(strings.TrimSpace(line))
			rules[ruleColors[0]] = ruleColors[1:]
		}

		if err != nil {
			break
		}
	}



	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return rules
}

//
// extract colors using a regex
//
func extractColors(data string)[]string {
	// [0-9]*\ *([a-z ]+)\ (bag)s*
	var colorsPattern = regexp.MustCompile("[0-9]*\\ *([a-z ]+)\\ (bag)s*")
	groups := colorsPattern.FindAllStringSubmatch(data, -1)

	result := []string{}

	for _, match := range groups {
		if (len(match) > 1) {
			result = append(result, match[1])
		}
	}

	return result
}

//
// recursive directed graph traversal to find all valid nodes reachable from color input
// leafs. For the puzzle starting with "shiny gold"
// 
func countRules(rules map[string][]string, markedColors map[string]bool, color string) int {
	var count int
	for rule, content := range rules {
		if !markedColors[rule] && isInColors(content, color) {
			markedColors[rule] = true
			count += countRules(rules, markedColors, rule) +1
		}
	}

	return count
}

//
// check if the passed color is in the list of given colors
// 
func isInColors(colors []string, color string) bool {
	for _, colorFromList := range colors {
		if (colorFromList == color) {
			return true
		}
	}

	return false
}
