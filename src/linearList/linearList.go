package main

import (
	"fmt"
)

type Element int

type LinearList struct {
	MaxSize int
	Length  int
	Data    []Element
}

//init
func New(maxsize int) *LinearList {
	return &LinearList{MaxSize: maxsize, Data: make([]Element, maxsize)}
}

func (list *LinearList) IsEmpty() bool {
	return 0 == list.Length
}

func (list *LinearList) IsFull() bool {
	return list.Length == list.MaxSize
}

//insert O(n)
func (list *LinearList) Insert(i int, e Element) bool {
	if i < 1 || i > list.Length {
		fmt.Println("pls check i:", i)
		return false
	}

	for k := list.Length; k > i-1; k-- {
		list.Data[k] = list.Data[k-1]
	}
	list.Data[i-1] = e
	list.Length++
	return true
}

//delete O(n)
func (list *LinearList) Del(i int) bool {
	if i < 1 || i > list.Length {
		fmt.Println("pls check i:", i)
		return false
	}
	for k := i - 1; k < list.Length-1; k++ {
		list.Data[k] = list.Data[k+1]
	}

	list.Data[list.Length-1] = 0
	list.Length--
	return true
}

//get O(1)
func (list LinearList) GetElem(i int) Element {
	if i < 1 || i > list.Length {
		fmt.Println("pls check i:", i)
		return -1
	}
	return list.Data[i-1]
}

//append O(1)
func (list *LinearList) append(e Element) bool {
	if list.IsFull() {
		fmt.Println("list is fulle")
		return false
	}
	list.Data[list.Length] = e
	list.Length++
	return true
}

func main() {
	ll := New(10)

	ll.append(99)
	ll.append(999)
	ll.append(9999)
	ll.append(99999)
	fmt.Println(ll)

	ll.Insert(4, 888)
	fmt.Println(ll)

	ll.Del(2)
	fmt.Println(ll)

	fmt.Println(ll.GetElem(3))
}
