package main

import (
	"crypto/sha256"
	"log"

	"../peachestree"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other peachestree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func main() {
	//Build list of Content to build tree
	var list []peachestree.Content
	list = append(list, TestContent{x: "Hello"})
	list = append(list, TestContent{x: "Hi"})
	list = append(list, TestContent{x: "Hey"})
	list = append(list, TestContent{x: "Hola"})

	//Create a new Merkle Tree from the list of Content
	tree, err := peachestree.NewTree(list, 3)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	rh := tree.RootHash()
	log.Println("rootHash: " + string(rh))

	//Verify the entire tree (hashes for each node) is valid
	//验证树
	//vt, err := tree.VerifyTree()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	//vc, err := t.VerifyContent(TestContent{x: "world"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Verify Content: ", vc)

	//String representation
	log.Println(tree)
}
