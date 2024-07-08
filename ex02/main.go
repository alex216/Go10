package main

import (
	"fmt"
	"piscine"
)

type TreeNode = piscine.TreeNode

func main() {
	root := &piscine.TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	piscine.BTreeApplyPreorder(root, fmt.Print)
	fmt.Println()

	root2 := &piscine.TreeNode{Data: "4"}
	BTreeInsertData(root2, "2")
	BTreeInsertData(root2, "6")
	BTreeInsertData(root2, "1")
	BTreeInsertData(root2, "3")
	BTreeInsertData(root2, "5")
	BTreeInsertData(root2, "7")
	piscine.BTreeApplyPreorder(root2, fmt.Print)
}

// helper
func recursiveInsert(root *TreeNode, data string, parent *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data, Parent: parent}
	}
	if data < root.Data {
		root.Left = recursiveInsert(root.Left, data, root)
	} else {
		root.Right = recursiveInsert(root.Right, data, root)
	}
	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	return recursiveInsert(root, data, nil)
}
