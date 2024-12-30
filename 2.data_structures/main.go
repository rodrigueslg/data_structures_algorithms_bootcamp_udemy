package main

import (
	"fmt"

	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/arrays"
	"github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/maps"
	mySlices "github.com/rodrigueslg/data_structures_algorithms_bootcamp_udemy/arrays/1/slices"
)

func main() {
	// runArrays()
	//runSlices()
	runMaps()
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
