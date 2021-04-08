package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(1, 2, 3)
	expected := []int{3, 2, 1}
	for i, value := range expected {
		listValue, _ := actual.Search(i)
		num := listValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestLinkedList_Insert(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList()
	actual.Insert(1)
	actual.Insert(2)
	actual.Insert(3)
	expected := []int{3, 2, 1}
	for i, value := range expected {
		listValue, _ := actual.Search(i)
		num := listValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestLinkedList_Search(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(1, 2, 3)
	expected := 2
	listValue, _ := actual.Search(1)
	assert.Equal(expected, listValue, fmt.Sprintf("%d and %d not equal, but expected", listValue, expected))
}

func TestLinkedList_Delete(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(3, 2, 1, 0)
	expected := []int{0, 1, 3}
	actual.Delete(2)
	for i, value := range expected {
		listValue, _ := actual.Search(i)
		num := listValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestLinkedList_Deletion(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(3, 2, 1, 0)
	expected := []int{1, 2, 3}
	actual.Deletion()
	for i, value := range expected {
		listValue, _ := actual.Search(i)
		num := listValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(5, 8, 7, 1, 4)
	expected := []int{1, 4, 5, 7, 8}
	actual.Sort()
	for i, value := range expected {
		listValue, _ := actual.Search(i)
		num := listValue.(int)
		assert.Equal(value, num, fmt.Sprintf("%d and %d not equal, but expected", num, value))
	}
}

func TestLinkedList_SearchError(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(5, 8, 7, 1, 4)
	_, err := actual.Search(5)
	assert.EqualError(err, fmt.Sprint("wrong id"))
}

func TestLinkedList_DeleteError(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList(5, 8, 7, 1, 4)
	err := actual.Delete(5)
	assert.EqualError(err, fmt.Sprint("wrong id"))
}

func TestLinkedList_SortError(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList('2', '1')
	err := actual.Sort()
	assert.EqualError(err, fmt.Sprint("wrong type"))
}
