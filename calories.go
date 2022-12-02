package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type ItemsCalories struct {
	values []int
}

func ParseCaloriesInput() []ItemsCalories {
	var inputSequence []ItemsCalories

	source, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	scanner := bufio.NewScanner(source)

	input := ItemsCalories{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputSequence = append(inputSequence, input)
			input = ItemsCalories{}
			continue
		}

		value, _ := strconv.Atoi(line)
		input.values = append(input.values, value)
	}

	return inputSequence
}

func (ic *ItemsCalories) calculate() int {
	total := 0

	for _, calories := range ic.values {
		total += calories
	}

	return total
}
