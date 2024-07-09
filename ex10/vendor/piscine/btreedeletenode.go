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
