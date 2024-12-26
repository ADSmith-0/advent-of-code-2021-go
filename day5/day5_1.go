package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day5_1() {
	file, err := os.Open("./day5/data-test.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid Grid

	for scanner.Scan() {
		text := scanner.Text()
		isValid, points := getPoints(text)
		if isValid {
			for _, point := range points {
				fmt.Println(point)
				grid[point.y][point.x]++
			}
		}
	}

	printGrid(&grid)

	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] >= 2 {
				count++
			}
		}
	}

	fmt.Printf("\nCount: %d\n\n", count)
}
