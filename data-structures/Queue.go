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

func (q *Queue) Dequeue() int {
	if q.head == nil {
		panic("empty queue")
	}

	val := q.head.value
	q.head = q.head.next
	q.length--

	if q.head == nil {
		q.tail = nil
	}

	return val
}

func (q *Queue) Peek() int {
	if q.head == nil {
		panic("empty queue")
	}
	return q.head.value
}

func main() {
	q := Queue{}

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	for node := q.head; node != nil; node = node.next {
		if node.next != nil {
			fmt.Printf("[ %d ] --> ", node.value)
		} else {
			fmt.Printf("[ %d ]", node.value)
		}
	}

}
