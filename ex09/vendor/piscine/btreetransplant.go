package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if node.Parent == nil {
		return rplc
	}
	if node.Parent.Left == node {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
