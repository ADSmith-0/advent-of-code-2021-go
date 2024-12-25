package day4

import (
	"log"
	"strconv"
	"strings"
)

func processCallouts(s string) []int {
	strs := strings.Split(s, ",")
	callouts := make([]int, 0, len(strs))

	for _, value := range strs {
		numVal, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Fatal("Could not convert from string to int", err)
		}

		callouts = append(callouts, int(numVal))
	}

	return callouts
}

func createSquares(s string) []Square {
	squares := make([]Square, 0, 5)
	num := -1

	for i, char := range strings.Split(s, "") {
		if char == " " {
			if num > -1 {
				squares = append(squares, Square{value: num, marked: false})
				num = -1
			}
			continue
		}

		convertedNum, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal("Could not convert string to int", err)
		}

		if num == -1 {
			num = convertedNum
		} else {
			num = (num * 10) + convertedNum
		}

		if i == len(s)-1 && num > -1 {
			squares = append(squares, Square{value: num, marked: false})
		}
	}

	return squares
}

func checkBingo(board *Board, coords Coordinates) bool {
	hBingo := true
	vBingo := true

	for i := 0; i < 5; i++ {
		if !(*board)[coords.y][i].marked {
			hBingo = false
		}
		if !(*board)[i][coords.x].marked {
			vBingo = false
		}
		if !hBingo && !vBingo {
			return false
		}
	}

	return true
}
