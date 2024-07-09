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

func TestBTreeDelete(t *testing.T) {
	tests := []struct {
		name            string   // test name
		input           []string // test case(btree)
		input_to_delete string   // test case(delete node)
		want            string   // expected result
	}{
		//	4
		//	2 6
		// 1 3 5 7
		{"葉なし", []string{"4", "2", "6", "1", "3", "5", "7"}, "1", "234567"},
		//	5
		//	3 o
		// 1 o
		{"左葉あり", []string{"5", "3", "1"}, "3", "15"},
		//	5
		//	3 o
		// o 4
		{"右葉あり", []string{"5", "3", "4"}, "3", "45"},
		//	5
		//	3 o
		// 1 4
		{"両葉あり", []string{"5", "3", "1", "4"}, "3", "145"},
		// 1
		{"葉なし、かつroot", []string{"1"}, "1", ""},
		//	5
		//	3 o
		// 1 o
		{"左葉あり、かつroot", []string{"5", "3", "1"}, "5", "13"},
		//	5
		//	o 7
		// o 6
		{"右葉あり、かつroot", []string{"5", "6", "7"}, "5", "67"},
		//	3
		// 1 4
		{"両葉あり、かつroot", []string{"3", "1", "4"}, "3", "14"},
	}

	// テストの実行
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := abstructTest(t, tc.input, tc.input_to_delete)
			if got != tc.want {
				t.Errorf("want[%q], got:[%q]", tc.want, got)
			}
		})
	}
}

// //////////////////////////////////////// helper functions
func abstructTest(t *testing.T, testarray []string, delete_target string) string {
	root := createBinaryTree(testarray)
	delete_node := piscine.BTreeSearchItem(root, delete_target)
	root = piscine.BTreeDeleteNode(root, delete_node)
	return checkResultFunc(t, root, delete_target)
}

func createBinaryTree(data []string) *TreeNode {
	root := &TreeNode{Data: data[0]}
	for _, str := range data[1:] {
		piscine.BTreeInsertData(root, str)
	}
	return root
}

func checkResultFunc(t *testing.T, root *TreeNode, data string) string {
	output := captureStdout(func() {
		piscine.BTreeApplyInorder(root, fmt.Print)
	})
	if piscine.BTreeSearchItem(root, data) != nil {
		t.Errorf("not deleted")
	}
	if piscine.BTreeIsBinary(root) == false {
		t.Errorf("tree is not binary")
	}
	return output
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
