package peachestree

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
)

type Content interface {
	CalculateHash() ([]byte, error)
	Equals(other Content) (bool, error)
}
type PeachTree struct {
	Root         *Node
	rootHash     []byte
	Leafs        []*Node
	hashStrategy func() hash.Hash
	multiNumber  int
}
type Node struct {
	Tree        *PeachTree
	Parent      *Node
	children    []*Node
	Hash        []byte
	IsLeaf      bool
	LeafContent Content
}

func NewTree(contents []Content, multiNumber int) (*PeachTree, error) {
	//设置hash策略
	var defaultHashStrategy = sha256.New
	tree := &PeachTree{
		hashStrategy: defaultHashStrategy,
		multiNumber:  multiNumber,
	}
	//根据内容创建叶子节点与根
	root, leafs, err := buildWithContent(contents, tree)
	if err != nil {
		return nil, err
	}
	tree.Root = root
	tree.Leafs = leafs
	tree.rootHash = root.Hash
	return tree, nil
}

func buildWithContent(contents []Content, tree *PeachTree) (*Node, []*Node, error) {
	if len(contents) == 0 {
		return nil, nil, errors.New("error: cannot construct tree with no content")
	}
	var leafs []*Node
	//遍历内容，一对一转换成叶节点
	for _, content := range contents {
		hash, err := content.CalculateHash()
		if err != nil {
			return nil, nil, err
		}
		leafs = append(leafs, &Node{
			LeafContent: content,
			IsLeaf:      true,
			Hash:        hash,
		})
	}
	//根据叶子节点递归创建中间节点
	root, err := buildWithIntermediate(leafs, tree)
	if err != nil {
		return nil, nil, err
	}
	return root, leafs, nil
}

func buildWithIntermediate(nodesToBuild []*Node, tree *PeachTree) (*Node, error) {
	var nodes []*Node
	multi := tree.multiNumber
	//循环遍历所有要构建的节点，按照multi为一组
	for i := 0; i < len(nodesToBuild); i += multi {
		h := tree.hashStrategy()
		n := &Node{
			Tree: tree,
		}
		//存放哈希和
		var allHash []byte
		var childrenList []*Node
		//j偏移量，要小于multi和剩下的个数
		for j := 0; j < (multi) && j < (len(nodesToBuild)-i); j++ {
			allHash = append(allHash, nodesToBuild[i+j].Hash...)
			childrenList = append(childrenList, nodesToBuild[i+j])
			nodesToBuild[i+j].Parent = n
		}
		//对哈希和进行二次哈希
		if _, err := h.Write(allHash); err != nil {
			return nil, err
		}
		n.children = childrenList
		n.Hash = h.Sum(nil)
		nodes = append(nodes, n)
		if len(nodesToBuild) <= multi {
			return n, nil
		}
	}
	return buildWithIntermediate(nodes, tree)
}

func (t *PeachTree) RootHash() []byte {
	return t.rootHash
}

//String returns a string representation of the node.
func (n *Node) String() string {
	return fmt.Sprintf("是否为叶子节点:%t\nHash:%v\n%s", n.IsLeaf, n.Hash, n.LeafContent)
}
func (m *PeachTree) String() string {
	s := ""
	for _, l := range m.Leafs {
		s += fmt.Sprint(l)
		s += "\n"
	}
	return s
}
