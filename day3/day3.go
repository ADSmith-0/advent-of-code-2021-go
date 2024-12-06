package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Main() {
	file, err := os.Open("./day3/data.txt")
	if err != nil {
		log.Fatal("Cannot open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bits [12]int
	numOfLines := 0

	for scanner.Scan() {
		numOfLines++
		str := scanner.Text()
		for index, value := range strings.Split(str, "") {
			if value == "1" {
				bits[index] = bits[index] + 1
			}
		}
	}

	gammaBits := make([]string, 0, 12)
	epsilsonBits := make([]string, 0, 12)

	for _, value := range bits {
		if value > (numOfLines / 2) {
			gammaBits = append(gammaBits, "1")
			epsilsonBits = append(epsilsonBits, "0")
		} else {
			gammaBits = append(gammaBits, "0")
			epsilsonBits = append(epsilsonBits, "1")
		}
	}

	gamma, err := strconv.ParseInt(strings.Join(gammaBits, ""), 2, 64)
	if err != nil {
		log.Fatal("Could not convert gammaBits:", err)
	}

	epsilon, err := strconv.ParseInt(strings.Join(epsilsonBits, ""), 2, 64)
	if err != nil {
		log.Fatal("Could not convert epsilonBits:", err)
	}

	fmt.Printf("Gamma: %d, Epsilon: %d\n", gamma, epsilon)
	fmt.Printf("Total: %d\n", gamma*epsilon)
}
