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
	rules, tickets, ownTicket := readTickets()

	filteredTickets := filterInvalid(rules, tickets)

	fmt.Printf("len tickets %d, len filtered tickets: %d\n", len(tickets), len(filteredTickets))

	orderedRules := determineRulesOrder(rules, filteredTickets)

	departureProd := prodDepartureRules(orderedRules, ownTicket)

	fmt.Printf("Product of departure rules of own ticket: %d\n", departureProd)
}

//
// struct to define a ticket validity rule
//
type ruleDef struct {
	desc        string
	lowerFirst  int
	upperFirst  int
	lowerSecond int
	upperSecond int
}

//
// read values from input and return as int slices
//
func readTickets() ([]ruleDef, [][]int, []int) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil, nil, nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	tickets := [][]int{}
	rules := []ruleDef{}
	yourTicket := []int{}
	inOtherTickets := false
	inYourTicket := false

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(strings.TrimSpace(line)) > 0 {
			if strings.Contains(line, ":") && strings.Contains(line, "or") {
				// rule
				rangeTokens := strings.Split(strings.Split(line, ":")[1], "or")
				var ruleElem ruleDef

				ruleElem.desc = strings.TrimSpace(line)

				rangeElem := strings.Split(rangeTokens[0], "-")
				ruleElem.lowerFirst, _ = strconv.Atoi(strings.TrimSpace(rangeElem[0]))
				ruleElem.upperFirst, _ = strconv.Atoi(strings.TrimSpace(rangeElem[1]))

				rangeElem = strings.Split(rangeTokens[1], "-")
				ruleElem.lowerSecond, _ = strconv.Atoi(strings.TrimSpace(rangeElem[0]))
				ruleElem.upperSecond, _ = strconv.Atoi(strings.TrimSpace(rangeElem[1]))

				rules = append(rules, ruleElem)
			} else if strings.Contains(line, "your ticket:") {
				inYourTicket = true
			} else if strings.Contains(line, "nearby tickets:") {
				inOtherTickets = true
				inYourTicket = false
			} else if inYourTicket {
				ticketTokens := strings.Split(line, ",")
				for _, ticketToken := range ticketTokens {
					number, _ := strconv.Atoi(strings.TrimSpace(ticketToken))
					yourTicket = append(yourTicket, number)
				}
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

	return rules, tickets, yourTicket
}

//
// filter out all invalid rules, rules with at least one number which doesn't fit
// in any rule
//
func filterInvalid(rules []ruleDef, tickets [][]int) [][]int {
	filteredTickets := [][]int{}

	for _, ticket := range tickets {
		var valid bool
		for _, val := range ticket {
			valid = false
			for _, rule := range rules {
				if isValid(rule, val) {
					valid = true
					break
				}

			}

			if !valid {
				break
			}
		}

		if valid {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	return filteredTickets
}

//
// determine the order of rules, under the assumption there's exactly
// one valid order of rules.
//
func determineRulesOrder(rules []ruleDef, tickets [][]int) []ruleDef {
	orderedRules := []ruleDef{}
	possibleRules := [][]ruleDef{}
	for i := 0; i < len(tickets[0]); i++ {
		for _, rule := range rules {
			valid := true
			for _, ticket := range tickets {
				val := ticket[i]

				if !isValid(rule, val) {
					valid = false
					break
				}
			}

			if valid {
				if len(possibleRules) < i+1 {
					possibleRules = append(possibleRules, []ruleDef{})
				}
				possibleRules[i] = append(possibleRules[i], rule)
			}
		}
	}

	orderedRules = make([]ruleDef, len(possibleRules))

	appliedRules := 0
	usedRules := make(map[string]bool)
	for appliedRules < len(orderedRules)-1 {
		for pos, rules := range possibleRules {

			if unusedRules := getUnusedRules(rules, usedRules); len(unusedRules) == 1 {
				orderedRules[pos] = unusedRules[0]
				usedRules[unusedRules[0].desc] = true
				appliedRules++
			}
		}
	}

	return orderedRules
}

//
// filter all unused rules from the list of rules
//
func getUnusedRules(rules []ruleDef, usedRules map[string]bool) []ruleDef {
	unusedRules := []ruleDef{}

	for _, rule := range rules {
		if !usedRules[rule.desc] {
			unusedRules = append(unusedRules, rule)
		}
	}

	return unusedRules
}

//
// check if the given value is within the ranges of the specified rule
//
func isValid(rule ruleDef, val int) bool {
	return (val >= rule.lowerFirst && val <= rule.upperFirst) ||
		(val >= rule.lowerSecond && val <= rule.upperSecond)
}

//
// calculate the product of all rules with the word "departure"
// in it from the values of the own ticket.
//
func prodDepartureRules(orderedRules []ruleDef, ticket []int) int {
	prod := 1

	for pos, rule := range orderedRules {
		if strings.Contains(rule.desc, "departure") {
			prod *= ticket[pos]
		}
	}

	return prod
}
