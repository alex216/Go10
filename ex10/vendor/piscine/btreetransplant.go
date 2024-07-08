package piscine

func BTreeTransplant(root, dst, rplc *TreeNode) *TreeNode {
	// error handling
	if root == nil || dst == nil {
		return root
	}

	// if dst is root
	if dst.Parent == nil {
		return rplc
	}

	// replace
	if dst.Parent.Left == dst {
		dst.Parent.Left = rplc
	} else {
		dst.Parent.Right = rplc
	}
	// update rplc.parent info
	if rplc != nil {
		rplc.Parent = dst.Parent
	}
	return root
}
