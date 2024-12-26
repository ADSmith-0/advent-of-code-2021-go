package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Grid [1000][1000]int

type Point struct {
	x int
	y int
}

func getPoints(s string) (bool, []Point) {
	strPoints := strings.Split(s, " -> ")
	points := make([]Point, 0, 10)

	for _, strCoord := range strPoints {
		splitPoints := strings.Split(strCoord, ",")
		numericX, err := strconv.ParseInt(splitPoints[0], 10, 32)
		if err != nil {
			log.Fatal("Could not convert string to int", err)
		}

		numericY, err := strconv.ParseInt(splitPoints[1], 10, 32)
		if err != nil {
			log.Fatal("Could not convert string to int", err)
		}

		points = append(points, Point{x: int(numericX), y: int(numericY)})
	}

	line := line(points[0], points[1])

	if line.dx != 0 && line.dy != 0 {
		return false, points
	}

	if line.dx == 0 {
		for y := line.y1 + line.deltaY; y != line.y2; y += line.deltaY {
			points = append(points, Point{x: line.x1, y: y})
		}
	}

	if line.dy == 0 {
		for x := line.x1 + line.deltaX; x != line.x2; x += line.deltaX {
			points = append(points, Point{x: x, y: line.y1})
		}
	}

	return true, points
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
