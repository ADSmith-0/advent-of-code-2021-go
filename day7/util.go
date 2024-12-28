package day7

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func getCrabPositions(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	crabPositions := make([]int, 0, 1000)

	num := 0

	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			crabPositions = append(crabPositions, num)
			num = 0
			continue
		}
		num = (num * 10) + digit
	}

	return crabPositions
}

func square(x int) int {
	return x * x
}

func calculateMinimalFuelCost(crabPositions *[]int, targetPosition int, stdDeviation int) int {
	smallestFuelCost := math.MaxInt32

	for i := targetPosition - stdDeviation*2; i <= targetPosition+stdDeviation*2; i++ {
		if i < 0 {
			continue
		}
		sum := 0
		for _, pos := range *crabPositions {
			sum += abs(pos - i)
		}
		if sum < smallestFuelCost {
			smallestFuelCost = sum
		}
	}

	return smallestFuelCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
