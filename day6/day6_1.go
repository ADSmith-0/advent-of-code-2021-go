package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day6_1() {
	file, err := os.Open("./day6/data.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	laternFish := make([]int, 0, 1000)

	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		laternFish = append(laternFish, num)
	}

	fmt.Printf("Initial state: ")
	printArr(&laternFish)

	for day := 1; day <= NUMBER_OF_DAYS; day++ {
		nextDay1(&laternFish)
	}

	fmt.Println("\nNumber of fish:", len(laternFish))
}

func nextDay1(laternFish *[]int) {
	for i := range *laternFish {
		if (*laternFish)[i] == 0 {
			(*laternFish)[i] = 6
			*laternFish = append(*laternFish, 8)
			continue
		}
		(*laternFish)[i] = (*laternFish)[i] - 1
	}
}
