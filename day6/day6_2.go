package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day6_2() {
	file, err := os.Open("./day6/data.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var laternFish [9]uint64

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		laternFish[num]++
	}

	for i := 1; i <= NUMBER_OF_DAYS; i++ {
		nextDay2(&laternFish)
	}

	var sum uint64 = 0

	for _, numberOfFish := range laternFish {
		sum += numberOfFish
	}

	fmt.Println("\nNumber of fish:", sum)

}

func nextDay2(laternFish *[9]uint64) {
	temp := (*laternFish)[0]
	(*laternFish)[0] = (*laternFish)[1]
	(*laternFish)[1] = (*laternFish)[2]
	(*laternFish)[2] = (*laternFish)[3]
	(*laternFish)[3] = (*laternFish)[4]
	(*laternFish)[4] = (*laternFish)[5]
	(*laternFish)[5] = (*laternFish)[6]
	(*laternFish)[6] = (*laternFish)[7]
	(*laternFish)[7] = (*laternFish)[8]
	(*laternFish)[8] = temp
	(*laternFish)[6] += temp
}
