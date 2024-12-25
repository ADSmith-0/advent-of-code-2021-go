package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day4_1() {
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
