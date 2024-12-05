package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	up Direction = iota
	down
	forward
)

type Movement struct {
	direction string
	amount    Direction
}

type CalculatePosition struct {
	addMovement func(string, int)
	getPosition func() (int, int)
}

func Main() {
	file, err := os.Open("./day2/data.txt")
	if err != nil {
		log.Fatal("Could not read file", err)
	}

	scanner := bufio.NewScanner(file)
	calcPos := calculatePosition()

	for scanner.Scan() {
		line := scanner.Text()
		lineValues := strings.Split(line, " ")
		num, err := strconv.Atoi(lineValues[1])
		if err != nil {
			log.Fatal("Could not convert to a number", err)
		}
		calcPos.addMovement(lineValues[0], num)
	}

	depth, position := calcPos.getPosition()
	fmt.Printf("depth: %d, position: %d\n", depth, position)
	fmt.Printf("total: %d\n", depth*position)
}

func calculatePosition() CalculatePosition {
	depth := 0
	position := 0

	addMovement := func(direction string, amount int) {
		switch direction {
		case "up":
			depth -= amount
		case "down":
			depth += amount
		case "forward":
			position += amount
		default:
			panic(fmt.Sprintf("unexpected day2.Direction: %#v", direction))
		}
	}

	getPosition := func() (int, int) {
		return depth, position
	}

	return CalculatePosition{addMovement: addMovement, getPosition: getPosition}
}
