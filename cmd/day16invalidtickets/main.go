// read input file input.txt and check for the first possible bus
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
// read in data, check the bus table for the earliest possible depart time and
// multiply bus id with waiting time
//
func main() {
	rules, tickets := readTickets()

	errorRate := checkValidity(rules, tickets)

	fmt.Printf("Error Rate of scanned tickets: %d\n", errorRate)
}

//
// struct to define a ticket validity rule
//
type rule struct {
	lowerFirst  int
	upperFirst  int
	lowerSecond int
	upperSecond int
}

//
// read values from input and return as int slices
//
func readTickets() ([]rule, [][]int) {
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
	tickets := [][]int{}
	rules := []rule{}
	inOtherTickets := false

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			if strings.Contains(line, ":") && strings.Contains(line, "or") {
				// rule
				rangeTokens := strings.Split(strings.Split(line, ":")[1], "or")
				var ruleElem rule

				rangeElem := strings.Split(rangeTokens[0], "-")
				ruleElem.lowerFirst, _ = strconv.Atoi(strings.TrimSpace(rangeElem[0]))
				ruleElem.upperFirst, _ = strconv.Atoi(strings.TrimSpace(rangeElem[1]))

				rangeElem = strings.Split(rangeTokens[1], "-")
				ruleElem.lowerSecond, _ = strconv.Atoi(strings.TrimSpace(rangeElem[0]))
				ruleElem.upperSecond, _ = strconv.Atoi(strings.TrimSpace(rangeElem[1]))

				rules = append(rules, ruleElem)
			} else if strings.Contains(line, "nearby tickets:") {
				inOtherTickets = true
			} else if inOtherTickets {
				ticketTokens := strings.Split(line, ",")
				ticket := []int{}
				for _, ticketToken := range ticketTokens {
					number, _ := strconv.Atoi(strings.TrimSpace(ticketToken))
					ticket = append(ticket, number)
				}
				tickets = append(tickets, ticket)
			}
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return rules, tickets
}

func checkValidity(rules []rule, tickets [][]int) int {
	var invalidTicketVals int

	for _, ticket := range tickets {
		for _, val := range ticket {
			valid := false
			for _, rule := range rules {
				if (val >= rule.lowerFirst && val <= rule.upperFirst) ||
					(val >= rule.lowerSecond && val <= rule.upperSecond) {
					valid = true
					break
				}

			}

			if !valid {
				invalidTicketVals += val
			}
		}
	}

	return invalidTicketVals
}
