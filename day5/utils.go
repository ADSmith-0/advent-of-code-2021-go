package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Grid [1000][1000]int

type Coordinates struct {
	x int
	y int
}

func getPoints(s string) (bool, []Coordinates) {
	strCoordinates := strings.Split(s, " -> ")
	rawCoordinates := make([][2]int, 0, 2)
	coordinates := make([]Coordinates, 0, 10)

	for _, strCoord := range strCoordinates {
		splitCoordinates := strings.Split(strCoord, ",")
		var numericCoordinates [2]int
		numericX, err := strconv.ParseInt(splitCoordinates[0], 10, 32)
		if err != nil {
			log.Fatal("Could not convert string to int", err)
		}
		numericCoordinates[0] = int(numericX)
		numericY, err := strconv.ParseInt(splitCoordinates[1], 10, 32)
		if err != nil {
			log.Fatal("Could not convert string to int", err)
		}
		numericCoordinates[1] = int(numericY)

		rawCoordinates = append(rawCoordinates, numericCoordinates)
	}

	startX := rawCoordinates[0][0]
	endX := rawCoordinates[1][0]
	startY := rawCoordinates[0][1]
	endY := rawCoordinates[1][1]
	rawXDiff := endX - startX
	rawYDiff := endY - startY

	if rawXDiff != 0 && rawYDiff != 0 {
		return false, coordinates
	}

	if rawXDiff == 0 {
		for yCoord := startY; yCoord != endY; yCoord += rawYDiff / abs(rawYDiff) {
			coordinates = append(coordinates, Coordinates{x: startX, y: yCoord})
		}
		coordinates = append(coordinates, Coordinates{x: startX, y: endY})
	}

	if rawYDiff == 0 {
		for xCoord := startX; xCoord != endX; xCoord += rawXDiff / abs(rawXDiff) {
			coordinates = append(coordinates, Coordinates{x: xCoord, y: startY})
		}
		coordinates = append(coordinates, Coordinates{x: endX, y: startY})
	}

	return true, coordinates
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func printGrid(grid *Grid) {
	fmt.Println()
	for y := range *grid {
		for x := range (*grid)[y] {
			if (*grid)[y][x] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", (*grid)[y][x])
			}
		}
		fmt.Println()
	}
}
