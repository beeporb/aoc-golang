package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

// A	Rock		1
// B 	Paper		2
// C 	Scissors	3

// X 	Rock		1
// Y 	Paper		2
// Z	Scissors	3

// 		Loss		0
//		Draw		3
//		Win			6
// Total Score = Val of Choice + Val of Outcome

const ROCK = 1
const PAPER = 2
const SCISSORS = 3

const LOSS = 0
const DRAW = 3
const WIN = 6

func matchChoice(choice *string) (int, error) {
	switch *choice {
	case "A":
		return ROCK, nil
	case "B":
		return PAPER, nil
	case "C":
		return SCISSORS, nil
	default:
		return -1, errors.New("invalid choice provided")
	}
}

func matchOutcome(outcome *string) (int, error) {
	switch *outcome {
	case "X":
		return LOSS, nil
	case "Y":
		return DRAW, nil
	case "Z":
		return WIN, nil
	default:
		return -1, errors.New("Invalid outcome provided.")
	}
}

func evaluateGame(game *[2]int) (int, error) {
	// Assumes 	game[1] == PLAYER
	// 			game[0] == OPPONENT
	// Returns player (game[1]) outcome score.
	switch *game {

	case [2]int{ROCK, WIN}:
		return WIN + PAPER, nil

	case [2]int{ROCK, DRAW}:
		return DRAW + ROCK, nil

	case [2]int{ROCK, LOSS}:
		return LOSS + SCISSORS, nil

	case [2]int{PAPER, WIN}:
		return WIN + SCISSORS, nil

	case [2]int{PAPER, DRAW}:
		return DRAW + PAPER, nil

	case [2]int{PAPER, LOSS}:
		return LOSS + ROCK, nil

	case [2]int{SCISSORS, WIN}:
		return WIN + ROCK, nil

	case [2]int{SCISSORS, DRAW}:
		return DRAW + SCISSORS, nil

	case [2]int{SCISSORS, LOSS}:
		return LOSS + PAPER, nil

	default:
		return -1, errors.New("invalid game state")
	}
}

func processLine(line string) int {
	line_components := strings.Split(line, " ")

	opponent, err := matchChoice(&line_components[0])

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	outcome, err := matchOutcome(&line_components[1])

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	game := [2]int{opponent, outcome}

	score, err := evaluateGame(&game)

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return score
}

func main() {

	fmt.Println("Opening game file...")

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0

	scores := make(chan int)

	wg := new(sync.WaitGroup)

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func() {
			score := processLine(line)
			scores <- score
		}()
	}

	go func() {
		for score := range scores {
			totalScore += score
			wg.Done()
		}
	}()

	wg.Wait()

	fmt.Println("All go routines finished.")

	fmt.Println("Total Score: ", totalScore)

}
