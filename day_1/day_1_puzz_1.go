package main

import (
	"bufio"
	"fmt"
	"os"
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
	maxElfCalories := 0

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
			if currentElfCalories > maxElfCalories {
				maxElfCalories = currentElfCalories
			}
			currentElfCalories = 0
		}
	}

	fmt.Println("Max Elf Calories: ", maxElfCalories)

}
