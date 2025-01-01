package main

import (
	"fmt"

	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/arrays"
	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/lists"
	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/maps"
	mySlices "github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/slices"
	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/stacksqueues"
)

func main() {
	// runArrays()
	//runSlices()
	//runMaps()
	//runLinkedLists()
	//runDoublyLinkedLists()
	//runStacks()
	runQueues()
}

func runArrays() {
	a := arrays.NewArray()
	a.Push("one")
	a.Push("two")
	a.Push("three")
	a.Push("four")

	// fmt.Println(a.Get(0))
	// fmt.Println(a.Get(1))
	// fmt.Println(a.Get(2))

	// a.Delete(1)
	// fmt.Println(a)

	// fmt.Println(a.Pop())
	// fmt.Println(a.Pop())

	// fmt.Println(a)
	// a.Delete(1)
	// fmt.Println(a)

	fmt.Println(a)
	a.Reverse()
	fmt.Println(a)
}

func runSlices() {
	s1 := []int{0, 3, 4, 31}
	s2 := []int{4, 6, 30}
	merged := mySlices.MergeSortedSlices(s1, s2)
	fmt.Println(merged)
}

func runMaps() {
	m := maps.NewMap(5)

	m.Set("grapes", 123)
	m.Set("grapes", "abc")
	m.Set("apples", true)
	m.Set("oranges", 12.5)

	fmt.Println(m.Get("grapes"))
	fmt.Println(m.Get("apples"))
	fmt.Println(m.Get("oranges"))

	fmt.Println(m.Keys())

	fmt.Println(m)
}

func runLinkedLists() {
	l := lists.NewLinkedList(10)
	l.Append(5)
	l.Append(16)
	l.Append(22)
	l.Append(45)
	l.Append(46)
	l.Append(48)

	l.Prepend(88)
	l.Insert(2, 24)

	l.Remove(0)
	l.Remove(2)

	fmt.Printf("List count: %d\n", l.Length)

	fmt.Println("List items:")
	array := l.ToArray()
	for i := range array {
		fmt.Println(array[i])
	}

	fmt.Println("Reversed items:")
	l.Reverse()
	reversedArray := l.ToArray()
	for i := range reversedArray {
		fmt.Println(reversedArray[i])
	}
}

func runDoublyLinkedLists() {
	l := lists.NewDoublyLinkedList(10)
	l.Append(5)
	l.Append(16)
	l.Append(22)
	l.Append(45)

	// l.Prepend(88)
	// l.Insert(2, 24)

	l.Remove(0)
	//l.Remove(2)

	fmt.Printf("List count: %d\n", l.Length)

	fmt.Println("List items:")
	node := l.Head
	for node != nil {
		fmt.Print("Prev: ")
		if node.Prev == nil {
			fmt.Print(" nil")
		} else {
			fmt.Print(node.Prev.Value)
		}

		fmt.Printf(", Value: %d", node.Value)

		fmt.Print(", Next: ")
		if node.Next == nil {
			fmt.Print(" nil")
		} else {
			fmt.Print(node.Next.Value)
		}

		fmt.Print("\n")

		node = node.Next
	}
}

func runStacks() {
	s := stacksqueues.NewStack()
	s.Push("google")
	s.Push("udemy")

	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func runQueues() {
	s := stacksqueues.NewQueue()
	s.Enqueue("Joy")
	s.Enqueue("Matt")
	s.Enqueue("Pavel")
	s.Enqueue("Samir")

	fmt.Println(s.Peek())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
}
