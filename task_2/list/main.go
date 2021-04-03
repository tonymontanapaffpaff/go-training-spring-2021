package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p *People) Compare(v1 interface{}, v2 interface{}) (int, error) {
	people1 := v1.(People)
	people2 := v2.(People)

	if people1.Age < people2.Age {
		return -1, nil
	}

	if people1.Age == people2.Age {
		return 0, nil
	}

	return 1, nil
}

func main() {
	l := New()

	p1 := People{
		Name: "Bill",
		Age:  18,
	}

	p2 := People{
		Name: "Ted",
		Age:  19,
	}

	l.Insert(p1, p2)

	l.Display()
	l.SortWithComparator(&People{})
	l.Display()

	fmt.Println("Searching result:")
	fmt.Println(l.Search(1))
	fmt.Println("Deleting by index result:")
	fmt.Println(l.Delete(1))
	fmt.Println("Deletion result:")
	fmt.Println(l.Deletion())

	l.Display()
}
