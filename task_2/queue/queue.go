package queue

import (
	"errors"
	"reflect"
)

//Node stores a value and a link to the next node in the queue
type Node struct {
	value interface{}
	next  *Node
}

// Implements a queue, stores a reference to the first and last item in the queue
// the queue has capacity and stores the current length
type Queue struct {
	len   int   // current length of the queue
	cap   int   // current capacity of the queue
	front *Node //Link to the first element of the queue
	back  *Node //Link to the last element of the queue
}

//Check if the queue is empty
// If the queue is empty returns true
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

//Check if the queue is full
// If the queue is full returns true
func (q *Queue) IsFull() bool {
	return q.len == q.cap
}

//  Add an element to the end of the queue
// Takes a value of type interface{}, returns error, when it occurs
func (q *Queue) Enqueue(value interface{}) error {
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
		return nil
	} else {
		return errors.New("queue is full")
	}
}

// Remove an element from the front of the queue
// returns element from the front of the queue and error, when it occurs
func (q *Queue) Dequeue() (interface{}, error) {
	if !q.IsEmpty() {
		elem := q.front
		q.front = q.front.next
		q.len--
		return elem.value, nil
	} else {
		return nil, errors.New("queue is empty")
	}
}

//Get the value of the front of the queue without removing it
// returns element from the front of the queue and error, when it occurs
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	} else {
		return q.front.value, nil
	}
}

//Sorts the list in ascending order, using the bubble method, returns error, when it occurs
func (q *Queue) Sort() error {
	elem := q.front
	for elem != nil {
		next := elem.next
		for next != nil {
			if comp, err := compare(elem.value, next.value); err == nil {
				if comp {
					elem.value, next.value = next.value, elem.value
				}
			} else {
				return err
			}
			next = next.next
		}
		elem = elem.next
	}
	return nil
}

//Comparing values for types: string, int float64
//Returns true if el1 is greater than el2
func compare(el1 interface{}, el2 interface{}) (bool, error) {
	if reflect.TypeOf(el1) != reflect.TypeOf(el2) {
		return false, errors.New("wrong type")
	}
	switch el1.(type) {
	case string:
		return el1.(string) > el2.(string), nil
	case int:
		return el1.(int) > el2.(int), nil
	case float64:
		return el1.(float64) > el2.(float64), nil
	default:
		return false, errors.New("wrong type")
	}
}

func NewQueue(cap int, values ...interface{}) (*Queue, error) {
	queue := new(Queue)
	queue.cap = cap
	for _, value := range values {
		if err := queue.Enqueue(value); err != nil {
			return queue, errors.New("queue is full")
		}
	}
	return queue, nil
}
