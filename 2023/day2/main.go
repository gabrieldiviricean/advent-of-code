package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id        int
	sets      []Cubes
	valid     bool
	minValues Cubes
	power     int
}

func (g *Game) isValid() {
	for _, set := range g.sets {
		if set.red > 12 || set.green > 13 || set.blue > 14 {
			g.valid = false
			return
		}
	}
	g.valid = true
}

func (g *Game) getMinValues() {
	reds := make([]int, 0)
	greens := make([]int, 0)
	blues := make([]int, 0)

	for _, set := range g.sets {
		if set.red != 0 {
			reds = append(reds, set.red)
		}
		if set.green != 0 {
			greens = append(greens, set.green)
		}
		if set.blue != 0 {
			blues = append(blues, set.blue)
		}
	}

	minRed := slices.Max(reds)
	minGreen := slices.Max(greens)
	minBlue := slices.Max(blues)

	g.minValues = Cubes{
		red:   minRed,
		green: minGreen,
		blue:  minBlue,
	}
}

func (g *Game) getPower() {
	power := g.minValues.red * g.minValues.green * g.minValues.blue
	g.power = power
}

func main() {
	fmt.Println("Day 2")

	// input, err := os.ReadFile("sample")
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
}

func partOne(input []byte) int {
	gamesList := strings.Split(string(input), "\n")
	games := make([]Game, 0)
	for _, game := range gamesList {
		gameInfo := strings.Split(game, ":")
		gameStr := strings.Split(strings.Trim(gameInfo[0], " "), " ")[1]
		gameId, _ := strconv.Atoi(gameStr)

		gameSets := strings.Split(gameInfo[1], ";")
		setsData := make([]Cubes, 0)
		gameData := Game{
			id:   gameId,
			sets: setsData,
		}
		for _, set := range gameSets {
			set = strings.Trim(set, " ")
			setCubes := strings.Split(set, ",")
			cubeData := Cubes{}
			for _, cube := range setCubes {
				cube = strings.Trim(cube, " ")
				cubeInfo := strings.Split(cube, " ")
				if cubeInfo[1] == "red" {
					cubeData.red, _ = strconv.Atoi(cubeInfo[0])
				} else if cubeInfo[1] == "green" {
					cubeData.green, _ = strconv.Atoi(cubeInfo[0])
				} else if cubeInfo[1] == "blue" {
					cubeData.blue, _ = strconv.Atoi(cubeInfo[0])
				}
			}
			gameData.sets = append(gameData.sets, cubeData)
			gameData.isValid()
		}
		games = append(games, gameData)
	}
	validIds := make([]int, 0)

	for _, game := range games {
		if game.valid {
			validIds = append(validIds, game.id)
		}
	}
	sum := 0
	for _, id := range validIds {
		sum += id
	}

	return sum
}

func partTwo(input []byte) int {
	gamesList := strings.Split(string(input), "\n")
	games := make([]Game, 0)
	for _, game := range gamesList {
		gameInfo := strings.Split(game, ":")
		gameStr := strings.Split(strings.Trim(gameInfo[0], " "), " ")[1]
		gameId, _ := strconv.Atoi(gameStr)

		gameSets := strings.Split(gameInfo[1], ";")
		setsData := make([]Cubes, 0)
		gameData := Game{
			id:   gameId,
			sets: setsData,
		}
		for _, set := range gameSets {
			set = strings.Trim(set, " ")
			setCubes := strings.Split(set, ",")
			cubeData := Cubes{}
			for _, cube := range setCubes {
				cube = strings.Trim(cube, " ")
				cubeInfo := strings.Split(cube, " ")
				if cubeInfo[1] == "red" {
					cubeData.red, _ = strconv.Atoi(cubeInfo[0])
				} else if cubeInfo[1] == "green" {
					cubeData.green, _ = strconv.Atoi(cubeInfo[0])
				} else if cubeInfo[1] == "blue" {
					cubeData.blue, _ = strconv.Atoi(cubeInfo[0])
				}
			}
			gameData.sets = append(gameData.sets, cubeData)
		}
		gameData.getMinValues()
		gameData.getPower()
		games = append(games, gameData)
	}

	sum := 0
	for _, game := range games {
		sum += game.power
	}

	return sum
}
