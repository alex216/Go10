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
	// root := &piscine.TreeNode{}

	node := piscine.BTreeSearchItem(root, "4")
	// node = nil

	replacement := &piscine.TreeNode{Data: "x"}
	// replacement := &piscine.TreeNode{}

	root = piscine.BTreeTransplant(root, node, replacement)
	piscine.BTreeApplyInorder(root, fmt.Println)
}
