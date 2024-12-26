package day5

type Line struct {
	x1     int
	x2     int
	y1     int
	y2     int
	dx     int
	dy     int
	adx    int
	ady    int
	deltaX int
	deltaY int
}

func line(p1 Point, p2 Point) Line {
	x1 := p1.x
	x2 := p2.x
	y1 := p1.y
	y2 := p2.y
	dx := x2 - x1
	dy := y2 - y1
	adx := abs(dx)
	ady := abs(dy)
	deltaX := 0
	deltaY := 0

	if adx > 0 {
		deltaX = dx / adx
	}

	if ady > 0 {
		deltaY = dy / ady
	}

	return Line{
		x1:     x1,
		x2:     x2,
		y1:     y1,
		y2:     y2,
		dx:     dx,
		dy:     dy,
		adx:    adx,
		ady:    ady,
		deltaX: deltaX,
		deltaY: deltaY,
	}
}
