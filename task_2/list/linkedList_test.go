package linkedList

import (
	"testing"
)

func TestLinkedListSize(t *testing.T) {
	l := New()

	l.Insert(1, 2, 3)

	if l.Size() != 3 {
		t.Errorf("The length isn't correct, expected: 3, found: %d\n", l.Size())
	}

	length := getLinkedListLength(l)
	if length != 3 {
		t.Errorf("The length of LinkedList isn't correct, expected: 3, found: %d\n", length)
	}

	// remove the element at the position 1
	v, err := l.Delete(1)
	if err != nil {
		t.Errorf("Failed to remove element with index 1, error: %v\n", err)
	}

	if v != 2 {
		t.Errorf("The removed element isn't correct, expect: 2, found: %v\n", v)
	}

	if l.Size() != 2 {
		t.Errorf("The length of LinkedList isn't correct, expected: 2 found: %d\n", l.Size())
	}

	length = getLinkedListLength(l)
	if length != 2 {
		t.Errorf("The length of LinkedList isn't correct, expected: 2 found: %d\n", length)
	}
}

func getLinkedListLength(ll List) int {
	length := 0
	it, hasNext := ll.Iterator()
	for hasNext {
		_, hasNext = it()
		length++
	}

	return length
}

func TestLinkedListSearch(t *testing.T) {
	l := New()

	l.Insert(1, 2, 3)

	v, err := l.Search(2)
	if err != nil {
		t.Errorf("Failed to get element at specific index, error: %v\n", err)
	}
	if v != 1 {
		t.Errorf("The element isn't expected, expected: 1, found: %v\n", v)
	}

	v, err = l.Search(100)
	if err == nil {
		t.Error("Expect error at searching element with index 100\n")
	}

	// check length at last
	if l.Size() != 3 {
		t.Errorf("The length isn't correct, expected: 3, found: %d\n", l.Size())
	}
}

func TestLinkedListIterator(t *testing.T) {
	l := New()

	l.Insert(1, 2, 3)

	it, hasNext := l.Iterator()
	if !hasNext {
		t.Error("The iterator should has elements")
	}

	// first element: 3
	v, hasNext := it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 3 {
		t.Errorf("The element isn't correct, expected: 3, found: %v\n", v)
	}

	// second element: 2
	v, hasNext = it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 2 {
		t.Errorf("The element isn't correct, expected: 2, found: %v\n", v)
	}

	// third element: 1
	v, hasNext = it()
	if hasNext {
		t.Error("The iterator shouldn't has next element")
	}
	if v != 1 {
		t.Errorf("The element isn't correct, expected: 1, found: %v\n", v)
	}
}

func TestLinkedListSort(t *testing.T) {
	l := New()

	l.Insert(4, 1, 3, 2)

	l.Sort()

	// check length after sorting
	if l.Size() != 4 {
		t.Errorf("The length isn't correct, expected: 4, found: %d\n", l.Size())
	}

	// check values after sorting
	if v, _ := l.Search(0); v != 1 {
		t.Errorf("The first element isn't correct, excepted: 1, found: %d\n", v)
	}
	if v, _ := l.Search(1); v != 2 {
		t.Errorf("The second element isn't correct, excepted: 2, found: %d\n", v)
	}
	if v, _ := l.Search(2); v != 3 {
		t.Errorf("The third element isn't correct, excepted: 3, found: %d\n", v)
	}
	if v, _ := l.Search(3); v != 4 {
		t.Errorf("The fourth element isn't correct, excepted: 4, found: %d\n", v)
	}
}

func TestLinkedListComparatorSort(t *testing.T) {
	l := New()

	l.Insert(
		&linkedListNode{age: 32}, &linkedListNode{age: 20},
		&linkedListNode{age: 27}, &linkedListNode{age: 25})

	l.SortWithComparator(&linkedListNode{})

	// check length after sorting
	if l.Size() != 4 {
		t.Errorf("The length isn't correct, expect: 4, found: %d\n", l.Size())
	}

	// check values after sorting
	if v, _ := l.Search(0); v.(*linkedListNode).age != 20 {
		t.Errorf("The first element isn't correct, excepted: 20, found: %d\n", v)
	}
	if v, _ := l.Search(1); v.(*linkedListNode).age != 25 {
		t.Errorf("The second element isn't correct, excepted: 25, found: %d\n", v)
	}
	if v, _ := l.Search(2); v.(*linkedListNode).age != 27 {
		t.Errorf("The third element isn't correct, excepted: 27, found: %d\n", v)
	}
	if v, _ := l.Search(3); v.(*linkedListNode).age != 32 {
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
