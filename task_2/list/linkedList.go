package linkedList

import (
	"errors"
	"fmt"

	"github.com/go-training-spring-2021/task_2/utils"
)

// List is a type of LinkedList.
type List interface {
	Insert(values ...interface{})
	Deletion() (interface{}, error)
	Display()
	Search(index int) (interface{}, error)
	Delete(index int) (interface{}, error)
	Iterator() (func() (interface{}, bool), bool)
	Size() int
	Sort()
	SortWithComparator(c utils.Comparator)
}

type element struct {
	next  *element
	value interface{}
}

// linkedList represents a singly linked list.
type linkedList struct {
	head   *element
	length int
}

// New initializes and returns an LinkedList.
func New() List {
	return &linkedList{
		head:   nil,
		length: 0,
	}
}

func (l *linkedList) size() int {
	return l.length
}

func (l *linkedList) isEmpty() bool {
	return l.size() == 0
}

func (l *linkedList) Insert(values ...interface{}) {
	for _, v := range values {
		l.linkBefore(v)
	}
}

// linkBefore inserts value before head element.
func (l *linkedList) linkBefore(value interface{}) {
	newElement := element{
		next:  l.head,
		value: value,
	}

	l.head = &newElement
	l.length++
}

// Deletion deletes an element at the beginning of the list.
func (l *linkedList) Deletion() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("empty list")
	}

	return l.unlink(l.head), nil
}

// Delete deletes an element using the id.
func (l *linkedList) Delete(index int) (interface{}, error) {
	size := l.size()

	if index < 0 || index >= size {
		return nil, fmt.Errorf("index out of range (len: %d, index: %d)", size, index)
	}

	return l.unlink(l.getElement(index)), nil
}

// unlink removes the specified element e in this list.
func (l *linkedList) unlink(e *element) interface{} {
	if e == nil {
		return nil
	}

	unlinkedValue := e.value

	if e.next != nil {
		e.value = e.next.value
		e.next = e.next.next
	} else {
		e.value = nil
	}

	l.length--
	return unlinkedValue
}

// getElement returns the element at the specified position.
func (l *linkedList) getElement(index int) *element {
	e := l.head

	for i := 0; i < index; i++ {
		e = e.next
	}

	return e
}

// Display displays the complete list.
func (l *linkedList) Display() {
	fmt.Println(l.values())
}

// Search searches an element using the id.
func (l *linkedList) Search(index int) (interface{}, error) {
	size := l.size()

	if index < 0 || index >= size {
		return nil, fmt.Errorf("index out of range (len: %d, index: %d)", size, index)
	}

	return l.getElement(index).value, nil
}

func (l *linkedList) Size() int {
	return l.length
}

// Sort sorts the element into ascending sequence using ReverseSort since the elements are inserted at the beginning
func (l *linkedList) Sort() {
	if l.size() < 2 {
		return
	}

	values := l.values()

	utils.ReverseSort(values, nil)

	l.clear()
	l.Insert(values...)
}

// SortWithComparator used when it is necessary to sort a list of structs
func (l *linkedList) SortWithComparator(c utils.Comparator) {
	if l.size() < 2 {
		return
	}

	values := l.values()

	utils.ReverseSort(values, c)

	l.clear()
	l.Insert(values...)
}

func (l *linkedList) clear() {
	for e := l.head; e != nil; {
		next := e.next
		e.next, e.value = nil, nil
		e = next
	}

	l.head, l.length = nil, 0
}

// values returns a list of values using an iterator
func (l *linkedList) values() []interface{} {
	if l.size() == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, l.size(), l.size())

	it, hasNext := l.Iterator()
	var v interface{}
	index := 0
	for hasNext {
		v, hasNext = it()
		values[index] = v
		index++
	}

	return values
}

// Iterator returns an iterator over the elements.
func (l *linkedList) Iterator() (func() (interface{}, bool), bool) {
	e := l.head

	return func() (interface{}, bool) {
		var element interface{}
		if e != nil {
			element = e.value
			e = e.next
		} else {
			element = nil
		}
		return element, e != nil
	}, e != nil
}
