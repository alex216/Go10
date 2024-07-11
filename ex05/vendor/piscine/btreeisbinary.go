package piscine

func rec(node *TreeNode, parent *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil && node.Left.Data > node.Data {
		return false
	}
	if node.Right != nil && node.Right.Data < node.Data {
		return false
	}
	if node.Parent != parent {
		return false
	}
	return (rec(node.Left, node) && rec(node.Right, node))
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Parent != nil {
		return false
	}
	return rec(root, nil)
}
