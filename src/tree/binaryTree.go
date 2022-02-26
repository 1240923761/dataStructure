package main

import "fmt"

type Element int

//定义二叉树的节点
type Node struct {
	Value Element
	Left  *Node
	Right *Node
}

//print node
func (node *Node) Print() {
	fmt.Printf("%d ", node.Value)
}

//set value
func (node *Node) SetValue(value Element) {
	node.Value = value
}

func CreateNode(value Element) *Node {
	return &Node{value, nil, nil}
}

//recursive
func (node *Node) FindNode(n *Node, x Element) *Node {
	if n == nil {
		return nil
	} else if n.Value == x {
		return n
	} else {
		p := node.FindNode(n.Left, x)
		if p != nil {
			return p
		}
		return node.FindNode(n.Right, x)
	}
}

//getHeight
func (node *Node) GetTreeHeight(n *Node) int {
	if n == nil {
		return 0
	} else {
		lHeigh := node.GetTreeHeight(n.Left)
		rHeigh := node.GetTreeHeight(n.Right)
		if lHeigh > rHeigh {
			return lHeigh + 1
		} else {
			return rHeigh + 1
		}
	}
}

//preOrder
func (node *Node) PreOrder(n *Node) {
	if n != nil {
		n.Print()
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

//功能：递归中序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) InOrder(n *Node) {
	if n != nil {
		node.InOrder(n.Left)
		n.Print()
		node.InOrder(n.Right)
	}
}

//功能：递归后序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) PostOrder(n *Node) {
	if n != nil {
		node.PostOrder(n.Left)
		node.PostOrder(n.Right)
		n.Print()
	}
}

//功能：打印所有的叶子节点
//参数：root
//返回值：nil
func (node *Node) GetLeafNode(n *Node) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			fmt.Printf("%d ", n.Value)
		}
		node.GetLeafNode(n.Left)
		node.GetLeafNode(n.Right)
	}
}

func main() {
	root := CreateNode(5)
	root.Left = CreateNode(2)
	root.Right = CreateNode(4)
	root.Left.Right = CreateNode(7)
	root.Left.Right.Left = CreateNode(6)
	root.Right.Left = CreateNode(8)
	root.Right.Right = CreateNode(9)

	fmt.Printf("%d\n", root.FindNode(root, 4).Value)
	fmt.Printf("%d\n", root.GetTreeHeight(root))

	root.PreOrder(root)
	fmt.Printf("\n")
	root.InOrder(root)
	fmt.Printf("\n")
	root.PostOrder(root)
	fmt.Printf("\n")

	root.GetLeafNode(root)
	fmt.Printf("\n")
}
