package main

import (
	"fmt"
	"piscine"
)

type Treenode = piscine.TreeNode

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeApplyByLevel(root, fmt.Print)
	fmt.Println()

	root2 := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root2, "2")
	piscine.BTreeInsertData(root2, "6")
	piscine.BTreeInsertData(root2, "1")
	piscine.BTreeInsertData(root2, "3")
	piscine.BTreeInsertData(root2, "5")
	piscine.BTreeInsertData(root2, "7")
	piscine.BTreeApplyByLevel(root2, fmt.Print)
}
