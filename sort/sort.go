package main

import (
	"fmt"
	"sort"
)

type Data struct {
	data float32
	aux  string
}

type MySort struct {
	data []Data
}

func (m *MySort) Len() int {
	return len(m.data)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (m *MySort) Less(i, j int) bool {
	return m.data[i].data < m.data[j].data
}

// Swap swaps the elements with indexes i and j.
func (m *MySort) Swap(i, j int) {
	tmp := m.data[i]
	m.data[i] = m.data[j]
	m.data[j] = tmp
}

func InsertionSort(data sort.Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && data.Less(j, j-1) {
			data.Swap(j, j-1)
			j--
		}
	}
}

func BubbleSort(data sort.Interface) {
	n := data.Len()
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func main() {
	l1 := []Data{
		{1, "one1"},
		{3, "three1"},
		{1, "one2"},
		{-3, "minusthree1"},
		{10, "ten"},
		{13, "thirteen"},
		{1.2532, "x"},
		{5.2382, "y"},
		{1, "one3"},
	}
	s1 := &MySort{l1}
	fmt.Println(l1)
	//sort.Sort(s1)
	//sort.Stable(s1)
	//InsertionSort(s1)
	BubbleSort(s1)
	fmt.Println(l1)
}
