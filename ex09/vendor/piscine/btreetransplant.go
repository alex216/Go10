package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	// null guard
	if root == nil || node == nil || rplc == nil{
		return root
	}

	// allows from rplc
	rplc.Parent = node.Parent
	rplc.Left = node.Left
	rplc.Right = node.Right

	// allows from rplc children
	if (rplc.Right != nil) {
		rplc.Right.Parent = rplc
	}
	if (rplc.Left != nil) {
		rplc.Left.Parent = rplc
	}

	// allows from Parent
	if node.Parent != nil {
		if node == node.Parent.Left {
			rplc.Parent.Left = rplc
		} else {
			rplc.Parent.Right = rplc
		}
	} else {
		// if node is root
		root = rplc
	}

	return root
}
