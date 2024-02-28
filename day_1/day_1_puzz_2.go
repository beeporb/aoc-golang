package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Opening file...")

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("An error occured when we opened the file...")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentElfCalories := 0
	var allElfCalories []int

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			line_val, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println("An error occured when parsing the line.")
				return
			}

			currentElfCalories += line_val
		} else {
			allElfCalories = append(allElfCalories, currentElfCalories)
			currentElfCalories = 0
		}
	}

	sort.Ints(allElfCalories)
	slices.Reverse(allElfCalories)

	topThreeElves := allElfCalories[:3]

	topThreeElfSum := 0

	for _, num := range topThreeElves {
		topThreeElfSum += num
	}

	fmt.Println("Top three elves are carrying: ", topThreeElfSum)

}
