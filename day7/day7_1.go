package day7

import (
	"fmt"
	"math"
)

func day7_1() {
	crabPositions := getCrabPositions("./day7/data.txt")

	sum := 0
	for _, pos := range crabPositions {
		sum += pos
	}
	mean := sum / len(crabPositions)

	stdDeviation := stdDev(&crabPositions, mean)

	stdDeviationSum := 0.0
	stdDeviationLength := 0.0
	for _, pos := range crabPositions {
		if pos >= (mean-stdDeviation*2) && pos <= (mean+stdDeviation*2) {
			stdDeviationSum += float64(pos)
			stdDeviationLength++
		}
	}
	stdDeviationMean := int(math.Round(stdDeviationSum / stdDeviationLength))

	fmt.Println("Total fuel cost:", calculateMinimalFuelCost(&crabPositions, stdDeviationMean, stdDeviation))
}

func stdDev(arr *[]int, mean int) int {
	deviationsSum := 0
	for _, pos := range *arr {
		deviationsSum += square(pos - mean)
	}
	stdDeviation := int(math.Round(math.Sqrt(float64(deviationsSum / len(*arr)))))

	return stdDeviation
}
