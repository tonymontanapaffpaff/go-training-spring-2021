package queue

import (
	"testing"
)

var minQueueLen = 16

func TestQueueUpdate(t *testing.T) {
	q := New(0)

	for i := 0; i < minQueueLen; i++ {
		q.Enqueue(i)
	}

	if !q.IsFull() {
		t.Errorf("Failed to fill queue, expected len: %v, found: %v\n", minQueueLen, q.Length())
	}

	for i := 0; i < 3; i++ {
		q.Dequeue()
		q.Enqueue(minQueueLen + i)
	}

	for i := 0; i < minQueueLen; i++ {
		v, err := q.Peek()
		if err != nil {
			t.Errorf("Failed to peek element, error: %v\n", err)
		}
		if v.(int) != i+3 {
			t.Errorf("Failed to peek element, expected: %v, found: %v\n", i+3, v)
		}

		q.Dequeue()
	}

	if !q.IsEmpty() {
		t.Errorf("Failed to clear queue, expected len: %v, found: %v\n", 0, q.Length())
	}
}

func TestQueueGet(t *testing.T) {
	q := New(0)

	for i := 0; i < minQueueLen; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < minQueueLen; i++ {
		v, err := q.Peek()

		if err != nil {
			t.Errorf("Failed to peek element, error: %v\n", err)
		}
		if v.(int) != i {
			t.Errorf("Failed to peek element, expected: %v, found: %v\n", i, v)
		}

		v, err = q.Dequeue()
		if err != nil {
			t.Errorf("Failed to dequeue element, error: %v\n", err)
		}
		if v != i {
			t.Errorf("Failed to deqeue element, expected: %v, found: %v\n", i, v)
		}
	}
}

func TestQueueLength(t *testing.T) {
	q := New(0)

	if q.Length() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
		if q.Length() != i+1 {
			t.Errorf("Failed to queue element, expected len: %v, found: %v\n", i+1, q.Length())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Dequeue()
		if q.Length() != 1000-i-1 {
			t.Errorf("Failed to dequeue element, expected len: %v, found: %v\n", 1000-i-1, q.Length())
		}
	}
}

func TestQueueSort(t *testing.T) {
	q := New(0)

	q.Enqueue("d")
	q.Enqueue("a")
	q.Enqueue("c")
	q.Enqueue("b")

	q.Sort()

	// check length after sorting
	if q.Length() != 4 {
		t.Errorf("The length isn't correct, expected: 4, found: %d\n", q.Length())
	}

	// check values after sorting
	if v, _ := q.Dequeue(); v != "a" {
		t.Errorf("The first element isn't correct, excepted: a, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v != "b" {
		t.Errorf("The second element isn't correct, excepted: b, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v != "c" {
		t.Errorf("The third element isn't correct, excepted: c, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v != "d" {
		t.Errorf("The fourth element isn't correct, excepted: d, found: %d\n", v)
	}
}

func TestLinkedListComparatorSort(t *testing.T) {
	q := New(0)

	q.Enqueue(&linkedListNode{age: 32})
	q.Enqueue(&linkedListNode{age: 20})
	q.Enqueue(&linkedListNode{age: 27})
	q.Enqueue(&linkedListNode{age: 25})

	q.SortWithComparator(&linkedListNode{})

	// check length after sorting
	if q.Length() != 4 {
		t.Errorf("The length isn't correct, expect: 4, found: %d\n", q.Length())
	}

	// check values after sorting
	if v, _ := q.Dequeue(); v.(*linkedListNode).age != 20 {
		t.Errorf("The first element isn't correct, excepted: 20, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v.(*linkedListNode).age != 25 {
		t.Errorf("The second element isn't correct, excepted: 25, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v.(*linkedListNode).age != 27 {
		t.Errorf("The third element isn't correct, excepted: 27, found: %d\n", v)
	}
	if v, _ := q.Dequeue(); v.(*linkedListNode).age != 32 {
		t.Errorf("The fourth element isn't correct, excepted: 32, found: %d\n", v)
	}
}

type linkedListNode struct {
	age int
}

func (aln *linkedListNode) Compare(v1, v2 interface{}) (int, error) {
	n1, n2 := v1.(*linkedListNode), v2.(*linkedListNode)

	if n1.age < n2.age {
		return -1, nil
	}

	if n1.age == n2.age {
		return 0, nil
	}

	return 1, nil
}
