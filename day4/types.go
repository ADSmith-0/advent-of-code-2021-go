package day4

type Square struct {
	value  int
	marked bool
}

type Coordinates struct {
	x int
	y int
}

type Board = [5][5]Square
