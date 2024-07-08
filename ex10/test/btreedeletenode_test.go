package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"piscine"
	"testing"
)

type TreeNode = piscine.TreeNode

//	4
//	2 6
// 1 3 5 7

func TestNoChildren(t *testing.T) {
	abstructTest(t, []string{"4", "2", "6", "1", "3", "5", "7"}, "1", "234567")
}

//	5
//	3 o
// 1 o

func TestLeftChildren(t *testing.T) {
	abstructTest(t, []string{"5", "3", "1"}, "3", "15")
}

//	5
//	3 o
// o 4

func TestRightChildren(t *testing.T) {
	abstructTest(t, []string{"5", "3", "4"}, "3", "45")
}

//	5
//	3 o
// 1 4

func TestBothChildren(t *testing.T) {
	abstructTest(t, []string{"5", "3", "1", "4"}, "3", "145")
}

// 1

func TestNoChildrenButRoot(t *testing.T) {
	abstructTest(t, []string{"1"}, "1", "")
}

//	5
//	3 o
// 1 o

func TestLeftChildrenButRoot(t *testing.T) {
	abstructTest(t, []string{"5", "3", "1"}, "5", "13")
}

//	5
//	o 7
// o 6

func TestRightChildrenButRoot(t *testing.T) {
	abstructTest(t, []string{"5", "6", "7"}, "5", "67")
}

//	3
// 1 4

func TestBothChildrenButRoot(t *testing.T) {
	abstructTest(t, []string{"3", "1", "4"}, "3", "14")
}

// //////////////////////////////////////// helper functions
func abstructTest(t *testing.T, testarray []string, delete_target string, expected string) {
	root := createBinaryTree(testarray)
	delete_node := piscine.BTreeSearchItem(root, delete_target)
	root = piscine.BTreeDeleteNode(root, delete_node)
	checkResultFunc(t, root, delete_target, expected)
}

func createBinaryTree(data []string) *TreeNode {
	root := &TreeNode{Data: data[0]}
	for _, str := range data[1:] {
		piscine.BTreeInsertData(root, str)
	}
	return root
}

func checkResultFunc(t *testing.T, root *TreeNode, data string, expected string) {
	output := captureStdout(func() {
		piscine.BTreeApplyInorder(root, fmt.Print)
	})
	// fmt.Printf("[%s]\n", output)
	if output != expected {
		t.Errorf("\n\ngot [%s] want [%s]\n\n", output, expected)
	}
	if piscine.BTreeSearchItem(root, data) != nil {
		t.Errorf("not deleted")
	}

	if piscine.BTreeIsBinary(root) == false {
		t.Errorf("tree is not binary")
	}
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
