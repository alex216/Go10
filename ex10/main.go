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
	piscine.BTreeInsertData(root, "8")
	piscine.BTreeInsertData(root, "6")
	piscine.BTreeInsertData(root, "9")
	piscine.BTreeInsertData(root, "2")
	node := piscine.BTreeSearchItem(root, "4")

	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	VisualizeBTree(root, "", true)

	root = piscine.BTreeDeleteNode(root, node)

	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	VisualizeBTree(root, "", true)
}

const (
    blue = "\033[34m"
    green = "\033[32m"
    pink  = "\033[35m"
    reset = "\033[0m"
)

func VisualizeBTree(root *piscine.TreeNode, prefix string, isLeft bool) {
    if root != nil {
        if root.Parent == nil {
            fmt.Printf("%s%s%s\n", green, root.Data, reset)
        } else {
            if isLeft {
                fmt.Printf("%s%s├── %s%s%s%s\n", green,  prefix, reset,  pink, root.Data, reset)
            } else {
                fmt.Printf("%s%s└── %s%s%s%s\n", green, prefix, reset, pink, root.Data, reset)
            }
        }

        newPrefix := prefix
        if root.Parent != nil {
            if isLeft {
                newPrefix += "│   "
            } else {
                newPrefix += "    "
            }
        }

        VisualizeBTree(root.Left, newPrefix, true)
        VisualizeBTree(root.Right, newPrefix, false)
    }
}
