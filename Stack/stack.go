package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}
func (s *Stack) Pop() int {
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}
func main() {
	fmt.Println("Stack Example")
}
