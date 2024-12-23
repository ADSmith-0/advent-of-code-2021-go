package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	value  int
	marked bool
}

type Coordinates struct {
	x int
	y int
}

type Board = [5][5]Square

func Main() {
	file, err := os.Open("./day4/data.txt")
	if err != nil {
		log.Fatal("File cannot be opened", err)
	}
	defer file.Close()

	boards := make([]Board, 0, 1000)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	callouts := processCallouts(scanner.Text())
	scanner.Scan() // Newline

	boardsIndex := 0
	var boardBuilder Board
	boardBuilderIndex := 0

	for scanner.Scan() {
		bingoRow := scanner.Text()
		if bingoRow == "" {
			boards = append(boards, boardBuilder)
			boardsIndex++
			boardBuilderIndex = 0
			continue
		}
		boardBuilder[boardBuilderIndex] = [5]Square(createSquares(bingoRow))
		boardBuilderIndex++
	}

	boards = append(boards, boardBuilder)

	var winningBoardIndex int
	var winningCalloutIndex int

	for calloutIndex, callout := range callouts {
		for boardIndex, board := range boards {
			for y := range board {
				for x, square := range board[y] {
					if !square.marked && square.value == callout {
						boards[boardIndex][y][x].marked = true
						if calloutIndex > 4 {
							if checkBingo((*[5][5]Square)(&boards[boardIndex]), Coordinates{x: x, y: y}) {
								winningBoardIndex = boardIndex
								winningCalloutIndex = calloutIndex
								goto WinningBoardFound
							}
						}
					}
				}
			}
		}
	}

WinningBoardFound:
	sum := 0

	for y := range boards[winningBoardIndex] {
		for _, square := range boards[winningBoardIndex][y] {
			if !square.marked {
				sum += square.value
			}
		}
	}

	fmt.Printf("sum: %d\ncallout: %d\ntotal: %d\n", sum, callouts[winningCalloutIndex], sum*callouts[winningCalloutIndex])

}

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
