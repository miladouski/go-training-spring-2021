package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 0, 1, 2, 3)
	expected := []int{0, 1, 2, 3}
	for _, value := range expected {
		queueValue, _ := actual.Dequeue()
		num := queueValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4)
	assert.Equal(true, actual.IsEmpty(), fmt.Sprintf("%t and %t not equal, but expected", actual.IsEmpty(), true))
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 0, 1, 2, 3)
	assert.Equal(true, actual.IsFull(), fmt.Sprintf("%t and %t not equal, but expected", actual.IsFull(), true))
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4)
	actual.Enqueue(0)
	actual.Enqueue(1)
	actual.Enqueue(2)
	actual.Enqueue(3)
	expected := []int{0, 1, 2, 3}
	for _, value := range expected {
		queueValue, _ := actual.Dequeue()
		assert.Equal(value, queueValue, fmt.Sprintf("%d and %d not equal, but expected", queueValue, value))
	}
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 0, 1, 2, 3)
	actual.Dequeue()
	expected := []int{1, 2, 3}
	for _, value := range expected {
		queueValue, _ := actual.Dequeue()
		assert.Equal(value, queueValue, fmt.Sprintf("%d and %d not equal, but expected", queueValue, value))
	}
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 1, 2, 3)
	num, _ := actual.Peek()
	expected := 1
	assert.Equal(expected, num.(int), fmt.Sprintf("%d and %d not equal, but expected", num.(int), expected))
}

func TestQueue_Sort(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 2, 0, 3, 1)
	actual.Sort()
	expected := []int{0, 1, 2, 3}
	for _, value := range expected {
		queueValue, _ := actual.Dequeue()
		assert.Equal(value, queueValue, fmt.Sprintf("%d and %d not equal, but expected", queueValue, value))
	}
}

func TestQueue_EnqueueError(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4, 0, 1, 2, 3)
	err := actual.Enqueue(4)
	assert.EqualError(err, fmt.Sprint("queue is full"))
}

func TestQueue_DequeueError(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4)
	_, err := actual.Dequeue()
	assert.EqualError(err, fmt.Sprint("queue is empty"))
}

func TestQueue_PeekError(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(4)
	_, err := actual.Peek()
	assert.EqualError(err, fmt.Sprint("queue is empty"))
}

func TestQueue_SortError(t *testing.T) {
	assert := assert.New(t)
	actual, _ := NewQueue(2, '2', '1')
	err := actual.Sort()
	assert.EqualError(err, fmt.Sprint("wrong type"))
}

func TestNewQueueError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewQueue(4, 0, 1, 2, 3, 4)
	assert.EqualError(err, fmt.Sprint("queue is full"))
}
