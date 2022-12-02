package main

import "fmt"

func main() {
	podium := Podium{0, 0, 0}
	for _, itemsCalories := range ParseCaloriesInput() {
		podium.handle(itemsCalories.calculate())
	}
	fmt.Printf("High value of calories input is: %d\n", podium.first)
	fmt.Printf("Total calories count from top 3 Elses is: %d\n", podium.sum())
}
