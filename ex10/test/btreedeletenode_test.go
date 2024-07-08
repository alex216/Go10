package test

import (
	"fmt"
	"piscine"
	"testing"
)

type TreeNode = piscine.TreeNode

func TestBTreeDeleteNode(t *testing.T) {
	root := &TreeNode{Data: "5"}
	piscine.BTreeInsertData(root, "3")
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "4")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "6")
	piscine.BTreeInsertData(root, "8")

	delete_target := root.Left.Left
	newRoot := piscine.BTreeDeleteNode(root, delete_target) // 1

	if piscine.BTreeSearchItem(newRoot, "1") != nil {
		piscine.BTreeApplyInorder(root, fmt.Print)
		t.Errorf("Node with value 1 was not deleted")
	}

	if piscine.BTreeIsBinary(newRoot) == false {
		t.Errorf("Tree is not binary")
	}
}
