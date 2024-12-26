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

	startPoint := getPoint(strPoints[0])
	endPoint := getPoint(strPoints[1])

	line := line(startPoint, endPoint)

	if line.dx != 0 && line.dy != 0 && line.adx != line.ady {
		return false, points
	}

	if line.adx == line.ady {
		x := line.x1
		y := line.y1

		for i := 0; i <= line.adx; i++ {
			points = append(points, Point{x: x, y: y})
			x += line.deltaX
			y += line.deltaY
		}
		return true, points
	}

	if line.dx == 0 {
		for y := line.y1; y != line.y2; y += line.deltaY {
			points = append(points, Point{x: line.x1, y: y})
		}
		points = append(points, Point{x: line.x1, y: line.y2})
		return true, points
	}

	if line.dy == 0 {
		for x := line.x1; x != line.x2; x += line.deltaX {
			points = append(points, Point{x: x, y: line.y1})
		}
		points = append(points, Point{x: line.x2, y: line.y1})
		return true, points
	}

	return true, points
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func getPoint(s string) Point {
	splitPoints := strings.Split(s, ",")

	numericX, err := strconv.ParseInt(splitPoints[0], 10, 32)
	if err != nil {
		log.Fatal("Could not convert string to int", err)
	}

	numericY, err := strconv.ParseInt(splitPoints[1], 10, 32)
	if err != nil {
		log.Fatal("Could not convert string to int", err)
	}

	return Point{x: int(numericX), y: int(numericY)}
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
