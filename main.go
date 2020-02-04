package main

func main() {
	r := NewGame()
	r.CreateWorld(20)

	s := NewSnake()
	s.Reset()
	r.SetSnake(s)

	r.Run()
	r.Exit()
}
