package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	len   int
	cap   int
	front *Node
	back  *Node
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) IsFull() bool {
	return q.len == q.cap
}

func (q *Queue) Enqueue(value interface{}) {
	if !q.IsFull() {
		if q.IsEmpty() {
			node := &Node{value: value}
			q.front = node
			q.back = node
		} else {
			back := &Node{value: value}
			q.back.next = back
			q.back = back
		}
		q.len++
	} else {
		panic("Queue is full")
	}
}

func (q *Queue) Dequeue() interface{} {
	if !q.IsEmpty() {
		elem := q.front
		q.front = q.front.next
		q.len--
		return elem.value
	} else {
		panic("Queue is empty")
	}
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty")
	} else {
		return q.front.value
	}
}

func (q *Queue) Sort() {
	elem := q.front
	for elem != nil {
		next := elem.next
		for next != nil {
			if compare(elem.value, next.value) {
				elem.value, next.value = next.value, elem.value
			}
			next = next.next
		}
		elem = elem.next
	}
}

func compare(el1 interface{}, el2 interface{}) bool {
	switch el1.(type) {
	case string:
		return el1.(string) > el2.(string)
	case int:
		return el1.(int) > el2.(int)
	case float64:
		return el1.(float64) > el2.(float64)
	default:
		panic("Wrong type")
	}
}

func main() {
	queue := Queue{cap: 10}
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(10)
	queue.Enqueue(9)
	queue.Enqueue(8)
	queue.Enqueue(7)
	queue.Enqueue(6)
	queue.Enqueue(5)
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Sort()
	fmt.Println(queue.IsFull())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Peek())
}
