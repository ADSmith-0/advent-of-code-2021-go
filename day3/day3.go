package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BINARY_LENGTH byte = 12

type Variable int

const (
	Oxygen Variable = iota
	CO2
)

type Frequency int

const (
	Most Frequency = iota
	Least
)

func Main() {
	file, err := os.Open("./day3/data.txt")
	if err != nil {
		log.Fatal("Cannot open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]string, 0, 1000)
	var bits [BINARY_LENGTH]int
	numOfLines := 0

	for scanner.Scan() {
		numOfLines++
		str := scanner.Text()
		data = append(data, str)
		for index, value := range strings.Split(str, "") {
			if value == "1" {
				bits[index] = bits[index] + 1
			}
		}
	}

	gammaBits := make([]byte, 0, BINARY_LENGTH)
	epsilsonBits := make([]byte, 0, BINARY_LENGTH)

	for _, value := range bits {
		if value >= (numOfLines / 2) {
			gammaBits = append(gammaBits, '1')
			epsilsonBits = append(epsilsonBits, '0')
		} else {
			gammaBits = append(gammaBits, '0')
			epsilsonBits = append(epsilsonBits, '1')
		}
	}

	gamma, err := strconv.ParseInt(string(gammaBits), 2, 64)
	if err != nil {
		log.Fatal("Could not convert gammaBits:", err)
	}

	epsilon, err := strconv.ParseInt(string(epsilsonBits), 2, 64)
	if err != nil {
		log.Fatal("Could not convert epsilonBits:", err)
	}

	fmt.Printf("Gamma: %d, Epsilon: %d\n", gamma, epsilon)
	fmt.Printf("Total: %d\n\n", gamma*epsilon)

	oxygenBinaryValue := findValue(Oxygen, data, 0)
	co2BinaryValue := findValue(CO2, data, 0)

	oxygenValue, err := strconv.ParseInt(oxygenBinaryValue, 2, 64)
	if err != nil {
		log.Fatal("Could not convert oxygenBinaryValue:", err)
	}

	co2Value, err := strconv.ParseInt(co2BinaryValue, 2, 64)
	if err != nil {
		log.Fatal("Could not convert co2BinaryValue:", err)
	}

	fmt.Printf("Oxygen value: %d, CO2 value: %d\n", oxygenValue, co2Value)
	fmt.Printf("Total: %d\n\n", oxygenValue*co2Value)
}

func findValue(variable Variable, data []string, index byte) string {
	var bit byte
	if variable == Oxygen {
		bit = commonBit(Most, data, index)
	}
	if variable == CO2 {
		bit = commonBit(Least, data, index)
	}
	dataCopy := make([]string, 0, len(data))
	for _, value := range data {
		if value[index] == bit {
			dataCopy = append(dataCopy, value)
		}
	}

	for _, value := range dataCopy {
		fmt.Println(index, string(bit), value)
	}
	fmt.Println("")

	if len(dataCopy) > 1 {
		return findValue(variable, dataCopy, index+1)
	}
	return dataCopy[0]
}

func commonBit(frequency Frequency, b []string, index byte) byte {
	count1 := 0
	count0 := 0

	for _, value := range b {
		if value[index] == '1' {
			count1++
		}
		if value[index] == '0' {
			count0++
		}
	}

	if frequency == Most {
		if count1 >= count0 {
			return '1'
		}
		return '0'
	}

	if count1 < count0 {
		return '1'
	}
	return '0'
}
