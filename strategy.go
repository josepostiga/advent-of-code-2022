package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type StrategyPlay struct {
	opponent string
	action   string
}

func ParseStrategy() []StrategyPlay {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var strategy []StrategyPlay

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		strategy = append(strategy, StrategyPlay{line[0], line[1]})
	}

	return strategy
}
