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
	check_array := [4]string{"4", "1", "7", "5"}
	for _, select_item := range check_array {

		fmt.Println()
		selected := piscine.BTreeSearchItem(root, select_item)
		fmt.Print("Item selected -> ")
		if selected != nil {
			fmt.Println(selected.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Parent of selected item -> ")
		if selected.Parent != nil {
			fmt.Println(selected.Parent.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Left child of selected item -> ")
		if selected.Left != nil {
			fmt.Println(selected.Left.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Right child of selected item -> ")
		if selected.Right != nil {
			fmt.Println(selected.Right.Data)
		} else {
			fmt.Println("nil")
		}
	}
}
