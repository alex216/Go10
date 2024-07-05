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
	fmt.Println(piscine.BTreeIsBinary(root))

	root1 := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root1, "1")
	piscine.BTreeInsertData(root1, "7")
	piscine.BTreeInsertData(root1, "5")
	root1.Left.Data = "6"
	fmt.Println(piscine.BTreeIsBinary(root1))
}
