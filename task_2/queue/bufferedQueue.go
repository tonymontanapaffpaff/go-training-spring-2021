package queue

import (
	"errors"

	"github.com/go-training-spring-2021/task_2/utils"
)

// Queue is a type of BufferedQueue.
type Queue interface {
	Enqueue(elem interface{})
	Dequeue() (interface{}, error)
	IsEmpty() bool
	IsFull() bool
	Peek() (interface{}, error)
	Length() int
	Capacity() int
	Sort()
	SortWithComparator(c utils.Comparator)
}

// BufferedQueue represents a single instance of the queue data structure.
type BufferedQueue struct {
	buffer            []interface{}
	head, tail, count int
}

// New initializes and returns an BufferedQueue.
func New(n int) Queue {
	// 16 is the smallest capacity that queue may have
	size := 16

	// n must be a power of 2 for bitwise modulus: x % n == x & (n - 1)
	if n > 0 && n&1 == 0 {
		size = n
	}

	return &BufferedQueue{
		buffer: make([]interface{}, size),
	}
}

// Enqueue adds an element to the end of the queue.
func (q *BufferedQueue) Enqueue(elem interface{}) {
	if q.IsFull() {
		q.resize()
	}

	q.buffer[q.tail] = elem

	// bitwise modulus
	q.tail = (q.tail + 1) & (len(q.buffer) - 1)
	q.count++
}

// Dequeue removes an element from the front of the queue.
func (q *BufferedQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}

	deqElem := q.buffer[q.head]

	q.buffer[q.head] = nil

	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buffer) - 1)
	q.count--

	// Resize down if buffer 1/4 full.
	if len(q.buffer) > 16 && (q.count<<2) == len(q.buffer) {
		q.resize()
	}
	return deqElem, nil
}

// Peek gets the value of the front of the queue without removing it
func (q *BufferedQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}

	return q.buffer[q.head], nil
}

// IsEmpty checks if the queue is empty
func (q *BufferedQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull checks if the queue is full
func (q *BufferedQueue) IsFull() bool {
	return q.Length() == q.Capacity()
}

func (q *BufferedQueue) Length() int {
	return q.count
}

func (q *BufferedQueue) Capacity() int {
	return len(q.buffer)
}

// Sort sorts the element into ascending sequence
func (q *BufferedQueue) Sort() {
	if q.count < 2 {
		return
	}

	utils.Sort(q.buffer[:q.count], nil)
}

// SortWithComparator used when it is necessary to sort a list of structs
func (q *BufferedQueue) SortWithComparator(c utils.Comparator) {
	if q.count < 2 {
		return
	}

	utils.Sort(q.buffer[:q.count], c)
}

// resize resizes the queue to fit exactly twice its current contents
// this can result in shrinking if the queue is less than half-full
func (q *BufferedQueue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buffer[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buffer[q.head:])
		copy(newBuf[n:], q.buffer[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buffer = newBuf
}
