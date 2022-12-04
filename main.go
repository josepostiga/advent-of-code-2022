package main

import "fmt"

func main() {
	var score int

	for _, round := range ParseStrategy() {
		result := Play(round)
		fmt.Printf("Strategy: Opponent plays '%s'; Outcome is '%s'. Score: %d\n", round.opponent, round.action, result.score)

		score += result.score
	}

	fmt.Printf("Final Score: %d\n", score)
}
