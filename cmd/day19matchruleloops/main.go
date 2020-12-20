// read input file input.txt and find all messages matching rule 42 and 31 combinations
// in contrast to first part the input data now contains two looping rules
// 8: 42 | 42 8 
// 11: 42 31 | 42 11 31
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
	messages, rules := readRulesAndMessages()

	results := getMessagesForRule(messages, rules, 0)

	fmt.Printf("The amout of messages complying with rule 0 are: %d\n", len(results))
}

type ruleDef struct {
	token string
	left  []int
	right []int
}

//
// read values from input and return calculations as string slices
//
func readRulesAndMessages() ([]string, map[int]ruleDef) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	messages := []string{}
	rules := map[int]ruleDef{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			if strings.Contains(trimmedLine, ":") {
				// rule
				tokens := strings.Split(trimmedLine, ":")
				rule := ruleDef{}
				ruleID, _ := strconv.Atoi(tokens[0])

				if strings.Contains(tokens[1], "|") {
					refTokens := strings.Split(tokens[1], "|")
					rule.left = extractRefs(strings.TrimSpace(refTokens[0]))
					rule.right = extractRefs(strings.TrimSpace(refTokens[1]))
				} else if strings.Contains(tokens[1], "\"") {
					rule.token = strings.TrimSpace(strings.ReplaceAll(tokens[1], "\"", ""))
				} else {
					rule.left = extractRefs(strings.TrimSpace(tokens[1]))
				}
				rules[ruleID] = rule
			} else {
				messages = append(messages, trimmedLine)
			}
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return messages, rules
}

//
// extract the refs from the string and create a list
// of int refs from it
//
func extractRefs(data string) []int {
	list := strings.Split(data, " ")
	result := []int{}

	for _, str := range list {
		if len(strings.TrimSpace(str)) > 0 {
			val, _ := strconv.Atoi(strings.TrimSpace(str))
			result = append(result, val)
		}
	}

	return result
}

//
// get all valid rules for the given rule set
//
func getMessagesForRule(messages []string, rules map[int]ruleDef, ruleID int) map[string]bool {
	result := make(map[string]bool)

	startSet := createStartSet()

	for _, message := range messages {

		for _, rule := range startSet {
			// 0 consist of the looping rules 8 and 11
			if applyRule(message, rule, rules, 0) == len(message) {
				result[message] = true
				break
			}
		}
	}

	return result
}

//
// create all possible starter rules generate from the following structure
//              0
//           /     \
//          8      11
//          |       |
//        42 8   42 11 31
// for the matching the top rules 0, 8, 11 are stripped away and ignored
// and instead a list of possible combinations of 42 and 31 rules are generated
// and used as starting rules
//
func createStartSet() []ruleDef {
	startSet := []ruleDef{}

	for i := 0; i < 10; i++ {
		left := ""
		for k := 0; k <= i; k++ {
			left = left + " 42"
		}
		for j := 0; j < 10; j++ {
			left = "42 " + left + " 31"

			rule := ruleDef{}
			rule.left = extractRefs(left)
			startSet = append(startSet, rule)
		}
	}

	return startSet

}

//
// try to apply the given rule to the message. Amount of matched tokens is returned
// if this is equal to the length of the message, the message matches the specified
// rule
//
func applyRule(message string, rule ruleDef, rules map[int]ruleDef, pos int) int {

	if pos >= len(message) {
		return 1
	}

	// check for token rule (leaf rule)
	if len(rule.token) > 0 {
		if rule.token == string(message[pos]) {
			return 1
		}
		return 0
	}

	// check for recursive (node) rule
	left := 0
	for _, ruleID := range rule.left {
		result := applyRule(message, rules[ruleID], rules, pos+left)
		if result > 0 {
			left += result
		} else {
			left = 0
			break
		}

	}

	right := 0
	if left == 0 {
		for _, ruleID := range rule.right {
			result := applyRule(message, rules[ruleID], rules, pos+right)
			if result > 0 {
				right += result
			} else {
				right = 0
				break
			}
		}
	}

	return left + right
}
