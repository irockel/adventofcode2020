// read input file input.txt and find all possible bags inside the shiny gold bag
// and sum them up.
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
// read in data and wander through the slices and check for the highest
// seat ID
//
func main() {
	rules := createRules()

	ruleCount := countRules(rules, "shiny gold")

	fmt.Printf("contained bags in shiny gold: %d\n", ruleCount-1 /* shiny gold bag itself is not! counted */)
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

		ruleColors := extractColors(strings.TrimSpace(line))
		if len(strings.TrimSpace(line)) > 0 {
			if !strings.Contains(line, "no other bags") {
				rules[ruleColors[0]] = ruleColors[1:]
			} else {
				rules[ruleColors[0]] = []string{}
			}
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
func extractColors(data string) []string {
	// [0-9]*\ *([a-z ]+)\ (bag)s*
	var colorsPattern = regexp.MustCompile("([0-9]*)\\ *([a-z ]+)\\ (bag)s*")
	groups := colorsPattern.FindAllStringSubmatch(data, -1)

	result := []string{}

	for _, match := range groups {
		if len(match) > 1 {
			result = append(result, match[1]+":"+match[2])
		}
	}

	return result
}

//
// recursive directed graph traversal to find all valid nodes reachable from color input
// leafs. For the puzzle starting with "shiny gold"
//
func countRules(rules map[string][]string, color string) int {
	count := 1
	rule := rules[":"+color]

	for _, token := range rule {
		tokens := strings.Split(token, ":")
		multiplier := 1
		if len(tokens[0]) > 0 {
			multiplier, _ = strconv.Atoi(tokens[0])
		}

		count += multiplier * countRules(rules, tokens[1])
	}

	return count
}
