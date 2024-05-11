package main

import "fmt"

//type Element int

type Node struct {
	Data Element
	Next *Node
}

type singlyLinkedList struct {
	Head   *Node
	Length int
}

// init create node
func CreateNode(v Element) *Node {
	return &Node{v, nil}
}

// create a nil linkedList
func CreateList() *singlyLinkedList {
	return &singlyLinkedList{CreateNode(0), 0}
}

// implementation
// i from 0 to length-1
func (list *singlyLinkedList) Insert(i int, v Element) {
	s := CreateNode(v)
	pre := list.Head
	for count := 0; count <= i; count++ {
		if count == i {
			s.Next = pre.Next
			pre.Next = s
			list.Length++
		}
		pre = pre.Next
	}
}

// delete
func (list *singlyLinkedList) Delete(i int) {
	//

	pre := list.Head
	for count := 0; count <= i; count++ {
		s := pre.Next
		if count == i {
			pre.Next = s.Next // s cannot be nil
			list.Length--
		}
		pre = pre.Next
	}
}

// 返回链表长度
func (list *singlyLinkedList) GetLength() int {
	return list.Length
}

// 查询值 v 所在的位置
func (list *singlyLinkedList) Search(v Element) int {
	pre := list.Head.Next
	for i := 1; i <= list.Length; i++ {
		if pre.Data == v {
			return i
		}
		pre = pre.Next
	}
	return 0
}

// 判空
func (list *singlyLinkedList) isNull() bool {
	pre := list.Head.Next
	if pre == nil {
		return true
	}
	return false
}

// print
func PrintList(list *singlyLinkedList) {
	pre := list.Head.Next
	fmt.Println("singlyLinkedList shows as follows: ...")
	for i := 1; i <= list.Length; i++ {
		fmt.Printf("%v\n", pre.Data)
		pre = pre.Next
	}
}

// main 函数：
func main() {
	singlyLinkedList := CreateList()
	fmt.Println("List is null: ", singlyLinkedList.isNull())

	singlyLinkedList.Insert(0, 3)
	singlyLinkedList.Insert(1, 6)
	singlyLinkedList.Insert(0, 5)

	PrintList(singlyLinkedList)
	fmt.Println("List length is: ", singlyLinkedList.Length)
	fmt.Println("元素6在位置：", singlyLinkedList.Search(6))
	fmt.Println("元素100在位置：", singlyLinkedList.Search(100))
	fmt.Println("List is null: ", singlyLinkedList.isNull())

	singlyLinkedList.Delete(2)
	PrintList(singlyLinkedList)
	fmt.Println("List length is: ", singlyLinkedList.Length)
}
