package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Main() {
	file, err := os.Open("./day1/data.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([]int, 0, 2000)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, num)
	}

	averageValues := make([]int, 0, 2000)
	initialAverageValue := 0

	for i := range values {
		if i <= 2 {
			initialAverageValue += values[i]
		}
		if i == 2 {
			averageValues = append(averageValues, initialAverageValue)
		}
		if i > 2 {
			averageValues = append(averageValues, averageValues[len(averageValues)-1]+values[i]-values[i-3])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for index, value := range averageValues {
		if index == 0 {
			continue
		}
		if value > averageValues[index-1] {
			count++
		}
	}

	fmt.Println("Number of increases", count)
}
