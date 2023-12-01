package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array []int) int {
	total := 0

	for _, n := range array {
		total += n
	}

	return total
}

func day1(rounds []string) int {
	roundScore := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	// A: rock: 1, B: paper: 2, C: scissors: 3
	// X: rock: 1, Y: paper: 2, Z: scissors: 3
	// loss: 0, draw: 3, win: 6

	allScores := []int{}
	for _, r := range rounds {
		score := roundScore[r]
		allScores = append(allScores, score)
	}

	total := sum(allScores)

	return total
}

func day2(rounds []string) int {
	roundScore := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}
	// A: rock: 1, B: paper: 2, C: scissors: 3
	// X: lose: 0, Y: draw: 3, Z: win: 6

	allScores := []int{}
	for _, r := range rounds {
		score := roundScore[r]
		allScores = append(allScores, score)
	}

	total := sum(allScores)

	return total
}

func main() {
	fmt.Println("Day 2")

	fmt.Println("Advent of code")

	input, err := os.ReadFile("input")
	check(err)

	rounds := strings.Split(string(input), "\n")
	day1Score := day1(rounds)
	fmt.Println(day1Score)
	day2Score := day2(rounds)
	fmt.Println(day2Score)

}
