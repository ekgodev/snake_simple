package main

type Point struct {
	X, Y float64
}

const (
	Top    = iota
	Right  = iota
	Bottom = iota
	Left   = iota
)

type Dir int

func (d Dir) Exec(p Point) Point {
	switch d {
	case Top:
		return Point{p.X, p.Y + 1}
	case Right:
		return Point{p.X + 1, p.Y}
	case Bottom:
		return Point{p.X, p.Y - 1}
	case Left:
		return Point{p.X - 1, p.Y}
	default:
		return p
	}
}

func (d Dir) CheckParallel(d2 Dir) bool {
	switch d {
	case Top:
		return d2 == Bottom
	case Right:
		return d2 == Left
	case Bottom:
		return d2 == Top
	case Left:
		return d2 == Right
	default:
		return false
	}
}
