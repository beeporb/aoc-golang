package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
	case "A", "X":
		return ROCK, nil
	case "B", "Y":
		return PAPER, nil
	case "C", "Z":
		return SCISSORS, nil
	default:
		return -1, errors.New("invalid choice provided")
	}
}

func evaluateGame(game *[2]int) (int, error) {
	// Assumes 	game[1] == PLAYER
	// 			game[0] == OPPONENT
	// Returns player (game[1]) outcome score.
	switch *game {

	case [2]int{ROCK, ROCK}:
		return DRAW + ROCK, nil

	case [2]int{ROCK, PAPER}:
		return WIN + PAPER, nil

	case [2]int{ROCK, SCISSORS}:
		return LOSS + SCISSORS, nil

	case [2]int{PAPER, PAPER}:
		return DRAW + PAPER, nil

	case [2]int{PAPER, ROCK}:
		return LOSS + ROCK, nil

	case [2]int{PAPER, SCISSORS}:
		return WIN + SCISSORS, nil

	case [2]int{SCISSORS, PAPER}:
		return LOSS + PAPER, nil

	case [2]int{SCISSORS, SCISSORS}:
		return DRAW + SCISSORS, nil

	case [2]int{SCISSORS, ROCK}:
		return WIN + ROCK, nil

	default:
		return -1, errors.New("invalid game state")
	}
}

func processLine(line string) int {
	line_components := strings.Split(line, " ")

	player, err := matchChoice(&line_components[0])

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	opponent, err := matchChoice(&line_components[1])

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	game := [2]int{player, opponent}

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

	// scores := make(chan int)

	// var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	score := processLine(line)
		// 	scores <- score
		// }()

		score := processLine(line)

		totalScore += score
	}

	// wg.Wait()

	fmt.Println("All go routines finished.")

	fmt.Println("Total Score: ", totalScore)

}
