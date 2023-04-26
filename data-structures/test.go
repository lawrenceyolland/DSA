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
	newNode := Node{value: val}

	q.length++
	if q.head == nil {
		q.head = &newNode
		q.tail = q.head
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
	return
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
		return 0, false
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
	fmt.Print("main")

	q := Queue{}

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	for node := q.head; node != nil; node = node.next {
		fmt.Printf("[ %d ]-->", node.value)
	}
}
