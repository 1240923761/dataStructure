package main

type Element int

type Node struct {
	Data Element
	Next *Node
}

type List struct {
	headNode *Node //头节点
}
