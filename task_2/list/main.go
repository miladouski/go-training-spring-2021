package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(value interface{}) {
	link := l.head
	l.head = &Node{value, link}
	l.len++
}

func (l *LinkedList) Deletion() {
	l.head = l.head.next
	l.len--
}

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

func (l *LinkedList) Search(id int) interface{} {
	if id < l.len {
		elem := l.head
		for i := 0; i < id; i++ {
			elem = elem.next
		}
		return elem.value
	} else {
		panic("wrong id")
	}
}

func (l *LinkedList) Delete(id int) {
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
	} else {
		panic("wrong id")
	}
}

func (l *LinkedList) Sort() {
	elem := l.head
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
	list := LinkedList{}
	list.Insert("0")
	list.Insert("1")
	list.Insert("2")
	list.Insert("3")
	list.Insert("4")
	list.Sort()
	list.Display()
	list.Delete(1)
	list.Display()
	list.Deletion()
	list.Display()
	fmt.Println(list.Search(1))
}
