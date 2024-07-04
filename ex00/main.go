package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	fmt.Println(root.Data)
	piscine.BTreeInsertData(root, "1")
	fmt.Printf(root.Left.Data)
	fmt.Println(root.Data)
	piscine.BTreeInsertData(root, "7")
	fmt.Printf(root.Left.Data)
	fmt.Printf(root.Data)
	fmt.Println(root.Right.Data)
	piscine.BTreeInsertData(root, "5")
	fmt.Printf(root.Left.Data)
	fmt.Printf(root.Data)
	fmt.Printf(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
}
