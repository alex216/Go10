package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeInsertData(root, "6") // toggle for test

	a := "1" // no children
	// a := "7" // no right children
	// a := "7"	// has both children when 6 is added
	// a := "4" // root
	node := piscine.BTreeSearchItem(root, a)

	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Printf("After delete %s:\n", a)
	piscine.BTreeApplyInorder(root, fmt.Println)
}
