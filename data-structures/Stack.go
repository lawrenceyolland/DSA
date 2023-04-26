package main

import "fmt"

type Node struct {
	value int
	prev  *Node
}

type Stack struct {
	head   *Node
	tail   *Node
	length int
}

func (s *Stack) Push(val int) {
	newNode := Node{value: val}
	s.length++

	if s.head == nil {
		s.head = &newNode
		return
	}

	newNode.prev = s.head
	s.head = &newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	val := s.head.value
	s.head = s.head.prev
	s.length--
	
	if s.head == nil {
		return 0, false
	}

	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	return s.head.value, true
}

func main() {
	s := Stack{}

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for node := s.head; node != nil; node = node.prev {
		if node.prev == nil {
			fmt.Printf("[ %d ]", node.value)
		} else {
			fmt.Printf("[ %d ]<--", node.value)
		}
	}
	// [ 4 ]<--[ 3 ]<--[ 2 ]<--[ 1 ]<--[ 0 ]

	s.Pop()
	fmt.Print("\n")

	for node := s.head; node != nil; node = node.prev {
		if node.prev == nil {
			fmt.Printf("[ %d ]", node.value)
		} else {
			fmt.Printf("[ %d ]<--", node.value)
		}
	}
	// [ 3 ]<--[ 2 ]<--[ 1 ]<--[ 0 ]  
}
