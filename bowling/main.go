package main

import (
	"fmt"
)

func calculateScore(rolls []int) int {
	score, frame, idx := 0, 0, 0

	for frame < 10 {
		if rolls[idx] == 10 {
			score += 10 + rolls[idx + 1] + rolls[idx + 2]
			idx++
		} else if rolls[idx] + rolls[idx + 1] == 10 {
			score += 10 + rolls[idx + 2]
			idx += 2
		} else {
			score += rolls[idx] + rolls[idx+1]
			idx += 2
		}

		frame++
	}

	return score
}

func main() {
	
	rolls2 := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10} 
	score2 := calculateScore(rolls2)
	fmt.Printf("Идеальная игра: %d\n", score2)

	
	rolls3 := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5} 
	score3 := calculateScore(rolls3)
	fmt.Printf("Все спары: %d\n", score3)
	
	rolls4 := []int{9, 0, 8, 1, 7, 2, 6, 3, 5, 4, 4, 5, 3, 6, 2, 7, 1, 8, 0, 9}
	score4 := calculateScore(rolls4)
	fmt.Printf("Только обычные фреймы: %d\n", score4)
	
	rolls5 := []int{10, 5, 5, 10, 5, 5, 10, 5, 5, 10, 5, 5, 10, 5, 5, 10}
	score5 := calculateScore(rolls5)
	fmt.Printf("Чередование страйков и спаров: %d\n", score5)

	rolls6 := []int{10, 7, 3, 9, 0, 10, 0, 8, 8, 2, 0, 6, 10, 10, 10, 8, 1}
	score6 := calculateScore(rolls6)
	fmt.Printf("Реальная игра: %d\n", score6)
}