// read input file input.txt and play combat until one player wins
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
// read in data, play the game and calculate the winning score
//
func main() {
	playerOne, playerTwo := readCardStacks()

	playerOne, playerTwo = playGame(playerOne, playerTwo)

	score := calculateWinnerScore(playerOne, playerTwo)

	fmt.Printf("The score of the winning player is : %d\n", score)
}

type ruleDef struct {
	token string
	left  []int
	right []int
}

//
// read values for the cards of the two players' stacks.
//
func readCardStacks() ([]int, []int) {
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
	playerOne := []int{}
	playerTwo := []int{}

	inPlayerOne := false
	inPlayerTwo := false

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			if strings.Contains(trimmedLine, "Player 1") {
				inPlayerOne = true
				inPlayerTwo = true
			} else if strings.Contains(trimmedLine, "Player 2") {
				inPlayerOne = false
				inPlayerTwo = true
			} else if inPlayerOne {
				card, _ := strconv.Atoi(trimmedLine)
				playerOne = append(playerOne, card)
			} else if inPlayerTwo {
				card, _ := strconv.Atoi(trimmedLine)
				playerTwo = append(playerTwo, card)
			}
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return playerOne, playerTwo
}

//
// play game of recursive combat
//
func playGame(playerOne, playerTwo []int) ([]int, []int) {
	stack := [2]int{}
	deckScores := make(map[int]map[int]bool)
	i := 0

	for len(playerOne) > 0 && len(playerTwo) > 0 && i < 1000 {
			// check if sub game needs to be played
		deckScoreA := calculateScore(playerOne)
		deckScoreB := calculateScore(playerTwo)
		

		if deckScores[deckScoreA][deckScoreB] {
			// avoid endless looping
			return playerOne, []int{}
		}

		if _, ok := deckScores[deckScoreA]; !ok {
			deckScores[deckScoreA] = map[int]bool{}
		}
		deckScores[deckScoreA][deckScoreB] = true

		stack[0] = playerOne[0]
		stack[1] = playerTwo[0]
		playerOne = playerOne[1:]
		playerTwo = playerTwo[1:]
	
		if stack[0] <= len(playerOne) && stack[1] <= len(playerTwo) {
			// recurse
			subA := playerOne
			if stack[0] < len(playerOne) {
				subA = make([]int, len(playerOne[:stack[0]]))
				copy(subA, playerOne[:stack[0]])
			}
			subB := playerTwo
			if stack[1] < len(playerTwo) {
				subB = make([]int, len(playerTwo[:stack[1]]))
				copy(subB, playerTwo[:stack[1]])
			}

			subA, _ = playGame(subA, subB)

			if len(subA) > 0 {
				// player one wins
				playerOne = append(playerOne, stack[0], stack[1])
			} else {
				// player two wins
				playerTwo = append(playerTwo, stack[1], stack[0])
			}

		} else {
			if stack[0] > stack[1] {
				// player one wins
				playerOne = append(playerOne, stack[0], stack[1])
			} else {
				// player two wins
				playerTwo = append(playerTwo, stack[1], stack[0])
			}
		}
		i++
	}


	return playerOne, playerTwo
}

//
// calculate winning score
//
func calculateWinnerScore(playerOne, playerTwo []int) int {
	var winner []int

	if len(playerOne) == 0 {
		winner = playerTwo
	} else {
		winner = playerOne
	}

	return calculateScore(winner)
}

//
// calculate score of the given deck
//
func calculateScore(deck []int) int {
	score := 0

	for i := 1; i <= len(deck); i++ {
		score += deck[len(deck)-i] * i
	}

	return score
}
