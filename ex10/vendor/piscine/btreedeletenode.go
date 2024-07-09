package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	var replace *TreeNode

	if node.Left == nil {
		replace = node.Right
	} else if node.Right == nil {
		replace = node.Left
	} else {
		replace = BTreeMin(node.Right)
		BTreeTransplant(root, replace, replace.Right)
		replace.Right = node.Right
		replace.Left = node.Left
	}

	if node == root {
		return replace
	}

	BTreeTransplant(root, node, replace)
	return root
}

// package piscine

// func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
// 	if node == nil {
// 		return root
// 	}
// 	if node.Left == nil {
// 		return BTreeTransplant(root, node, node.Right)
// 	} else if node.Right == nil {
// 		return BTreeTransplant(root, node, node.Left)
// 	} else {
// 		y := BTreeMin(node.Right)
// 		if y.Parent != node {
// 			BTreeTransplant(root, y, y.Right)
// 			y.Right = node.Right
// 			y.Right.Parent = y
// 		}
// 		BTreeTransplant(root, node, y)
// 		y.Left = node.Left
// 		y.Left.Parent = y
// 		return root
// 	}
// }
