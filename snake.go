package main

type Snake struct {
	Parts []Point
}

func NewSnake() *Snake {
	snake := &Snake{}
	return snake
}

func (s *Snake) Reset() {
	sx, sy, l := 1, 1, 7
	for i := l - 1; i >= 0; i-- {
		s.Parts = append(s.Parts, Point{float64(sx + i), float64(sy)})
	}
}

func (s *Snake) Len() int {
	return len(s.Parts)
}

func (s *Snake) Head() Point {
	if s.Len() == 0 {
		return Point{-1, -1}
	}
	return s.Parts[0]
}

func (s *Snake) Add(p Point) {
	s.Parts = append([]Point{p}, s.Parts...)
}

func (s *Snake) Tail() Point {
	if s.Len() == 0 {
		return Point{-1, -1}
	}
	return s.Parts[len(s.Parts)-1]
}

func (s *Snake) Move(d Dir) {
	lastP := s.Parts[0]
	s.Parts[0] = d.Exec(s.Parts[0])

	for i := range s.Parts[1:] {
		s.Parts[1+i], lastP = lastP, s.Parts[1+i]
	}
}

func (s *Snake) IsSnake(p Point) bool {
	for i := range s.Parts {
		if s.Parts[i] == p {
			return true
		}
	}
	return false
}

func (s *Snake) CutIfSnake(p Point) bool {
	i := 0
	for ; i < len(s.Parts); i++ {
		if s.Parts[i] == p {
			break
		}
	}

	s.Parts = s.Parts[0:i]
	return i >= len(s.Parts)
}
