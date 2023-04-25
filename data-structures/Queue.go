package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) Enqueue(val int) {
	q.length++

	newNode := Node{value: val}
	if q.tail == nil {
		q.head = &newNode
		q.tail = &newNode
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
}

func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	val := q.head.value
	q.head = q.head.next
	q.length--

	if q.head == nil {
		q.tail = nil
	}

	return val, true
}

func (q *Queue) Peek() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	return q.head.value, true
}

func main() {
	q := Queue{}

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	for node := q.head; node != nil; node = node.next {
		if node.next != nil {
			fmt.Printf("[ %d ] --> ", node.value)
		} else {
			fmt.Printf("[ %d ]", node.value)
		}
	}
	// [ 0 ] --> [ 1 ] --> [ 2 ] --> [ 3 ] --> [ 4 ]
}
