package piscine

import (
	_ "fmt"
)

type MyQueue struct {
	list *List
}

func NewQueue() *MyQueue {
	return &MyQueue{
		list: &List{},
	}
}

func (queue *MyQueue) deque() (*TreeNode, error) {
	node, err := ListPopFirst(queue.list)
	if node == nil || err != nil {
		return nil, err
	}
	return node.Data.(*TreeNode), nil
}

func (queue *MyQueue) enque(node *TreeNode) {
	ListPushBack(queue.list, node)
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := NewQueue()
	queue.enque(root)

	for {
		treeNode, err := queue.deque()
		if treeNode == nil || err != nil {
			break
		}
		f(treeNode.Data)
		if treeNode.Left != nil {
			queue.enque(treeNode.Left)
		}
		if treeNode.Right != nil {
			queue.enque(treeNode.Right)
		}
	}
	return
}
