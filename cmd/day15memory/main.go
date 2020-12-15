// read input file input.txt and check for the first possible bus
package main

import (
	"fmt"
)

//
// read in data, check the bus table for the earliest possible depart time and
// multiply bus id with waiting time
//
func main() {
	startNumbers := []int{0, 3, 1, 6, 7, 5}

	stopValue2020, stopValue30 := playMemory(startNumbers, 30000000)

	fmt.Printf("The stop value in the 2020 round is: %d\n", stopValue2020)
	fmt.Printf("The stop value in the 30000000 round is: %d\n", stopValue30)
}

//
// run the specified program
//
func playMemory(startNumbers []int, stopRound int) (int, int) {
	rounds := []int{}
	numberIndex := make(map[int]int)

	rounds = append(rounds, startNumbers...)

	for pos, val := range rounds {
		numberIndex[val] = pos
	}

	for i := len(rounds); i <= stopRound; i++ {
		lastVal := rounds[i-1]
		if index, ok := numberIndex[lastVal]; ok {
			newVal := i - 1 - index
			rounds = append(rounds, newVal)
		} else {
			rounds = append(rounds, 0)
		}
		numberIndex[lastVal] = i - 1
	}

	return rounds[2020-1], rounds[stopRound-1]
}
