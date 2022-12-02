package main

import "fmt"

func main() {
	high := 0
	for _, itemsCalories := range ParseCaloriesInput() {
		totalItemCalories := itemsCalories.calculate()
		if totalItemCalories > high {
			high = totalItemCalories
		}
	}
	fmt.Printf("Higher calories count is: %d\n", high)
}
