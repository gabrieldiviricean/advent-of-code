package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

func main() {
	fmt.Println("Day 1")

	input, err := os.ReadFile("day1/input")
	check(err)

	inputArray := strings.Split(string(input), "\r\n\r\n")

	highestCalories := 0
	allCalories := []int{}

	for _, s := range inputArray {
		caloriesArray := []int{}
		splitArray := strings.Split(s, "\r\n")

		for _, s := range splitArray {
			n, err := strconv.Atoi(s)
			check(err)

			caloriesArray = append(caloriesArray, n)
		}
		sum := sum(caloriesArray)

		if sum > highestCalories {
			highestCalories = sum
		}

		allCalories = append(allCalories, sum)

	}

	sort.Ints(allCalories)

	biggestThree := allCalories[len(allCalories)-3:]
	totalThree := sum(biggestThree)

	fmt.Println(highestCalories)
	fmt.Println(totalThree)

}
