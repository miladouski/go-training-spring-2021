package list

import (
	"errors"
	"fmt"
	"reflect"
)

//Node stores a value and a link to the next node in the list
type Node struct {
	value interface{} // Stores the value of the list
	next  *Node       //Link to next node
}

// Implementing a linked list, stores a link to the first element of the list
// and the current length
type LinkedList struct {
	head *Node //Link to the first element of the list
	len  int   //current length of the list
}

//Adds an element at the beginning of the list
//Takes a value of type interface{}
func (l *LinkedList) Insert(value interface{}) {
	link := l.head
	l.head = &Node{value, link}
	l.len++
}

// Deletes an element at the beginning of the list
func (l *LinkedList) Deletion() {
	l.head = l.head.next
	l.len--
}

// Displays the complete list
func (l *LinkedList) Display() {
	if l.len > 0 {
		elem := l.head
		for i := 0; i < l.len; i++ {
			fmt.Printf("%v ", elem.value)
			elem = elem.next
		}
	}
	fmt.Println()
}

//Searches an element using the id
// Takes id of an element in a list, returns a value from a list and error
func (l *LinkedList) Search(id int) (interface{}, error) {
	if id < l.len {
		elem := l.head
		for i := 0; i < id; i++ {
			elem = elem.next
		}
		return elem.value, nil
	} else {
		return nil, errors.New("wrong id")
	}
}

//Deletes an element using the id
// Takes id of an element in a list, returns error, when it occurs
func (l *LinkedList) Delete(id int) error {
	if id < l.len {
		if id == 0 {
			l.head = l.head.next
		} else {
			elem := l.head
			for i := 0; i < id-1; i++ {
				elem = elem.next
			}
			elem.next = elem.next.next
		}
		l.len--
		return nil
	} else {
		return errors.New("wrong id")
	}
}

//Sorts the list in ascending order, using the bubble method, returns error, when it occurs
func (l *LinkedList) Sort() error {
	elem := l.head
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

func NewLinkedList(values ...interface{}) *LinkedList {
	list := new(LinkedList)
	for _, node := range values {
		list.Insert(node)
	}
	return list
}
