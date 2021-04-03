package main

import "fmt"

func main() {
	q := New(0)

	q.Enqueue("c")
	q.Enqueue("a")
	q.Enqueue("b")

	fmt.Println("Peek before sorting:")
	fmt.Println(q.Peek())

	q.Sort()

	fmt.Println("Peek after sorting:")
	fmt.Println(q.Peek())

	fmt.Println("IsFull:", q.IsFull())
	fmt.Println("Dequeue result:")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println("IsEmpty:", q.IsEmpty())
}
