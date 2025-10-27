package main

import (
	"errors"
	"fmt"
)

// реализовать методы Pop(), Push(), IsEmpty(), Size(), Clear()

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("cтек пуст")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = nil
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Push(9)

	fmt.Println(s)

	s.Clear()
	fmt.Println("После очистки:", s)
	fmt.Println("Размер:", s.Size())
}
