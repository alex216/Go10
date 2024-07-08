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
		new_node := BTreeMin(node.Right)
		BTreeDeleteNode(root, new_node)
		BTreeTransplant(root, node, new_node)
	}

	if node == root {
		return replace
	}

	BTreeTransplant(root, node, replace)
	return root
}
