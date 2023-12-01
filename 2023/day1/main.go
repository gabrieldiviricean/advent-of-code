package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")

	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
}

func partOne(input []byte) int {
	allNumbers := make([]int, 0)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		number := findIntegers(line)
		allNumbers = append(allNumbers, number)
	}

	sum := 0
	for _, number := range allNumbers {
		sum += number
	}

	return sum
}

func partTwo(input []byte) int {
	allNumbers := make([]int, 0)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		line = replaceStringNumbers(line)
		number := findIntegers(line)
		allNumbers = append(allNumbers, number)
	}

	sum := 0
	for _, number := range allNumbers {
		sum += number
	}

	return sum
}

func findIntegers(s string) int {
	re := regexp.MustCompile(`\d`)
	numbers := re.FindAllString(s, -1)

	var value int

	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		value, _ = strconv.Atoi(numbers[0] + numbers[0])
	} else {
		value, _ = strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	}

	return value
}

func replaceStringNumbers(s string) string {
	s = strings.ReplaceAll(s, "one", "o1e")
	s = strings.ReplaceAll(s, "two", "t2o")
	s = strings.ReplaceAll(s, "three", "t3e")
	s = strings.ReplaceAll(s, "four", "f4r")
	s = strings.ReplaceAll(s, "five", "f5e")
	s = strings.ReplaceAll(s, "six", "s6x")
	s = strings.ReplaceAll(s, "seven", "s7n")
	s = strings.ReplaceAll(s, "eight", "e8t")
	s = strings.ReplaceAll(s, "nine", "n9e")

	return s
}
