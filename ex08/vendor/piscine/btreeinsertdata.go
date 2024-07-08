package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func recursiveInsert(root *TreeNode, data string, parent *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data, Parent: parent}
	}
	if data < root.Data {
		root.Left = recursiveInsert(root.Left, data, root)
	} else {
		root.Right = recursiveInsert(root.Right, data, root)
	}
	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	return recursiveInsert(root, data, nil)
}
