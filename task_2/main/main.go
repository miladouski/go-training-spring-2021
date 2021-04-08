package main

import (
	"fmt"
	"go-training-spring-2021/task_2/list"
	"go-training-spring-2021/task_2/queue"
)

func main() {
	queue, _ := queue.NewQueue(4, 0, 1, 2, 3)
	list := list.NewLinkedList(3, 2, 1, 0)
	list.Display()
	if value, err := queue.Peek(); err != nil {
		fmt.Println("error:", value)
	} else {
		fmt.Println(value)
	}
}
