package tree

import (
	"strconv"
	"testing"
)

type TreeNode int

func (t TreeNode) Less(e Element) bool {
	return t < e.(TreeNode)
}

func (t TreeNode) String() string {
	return strconv.Itoa(int(t))
}

func TestBinaryTree_Insert(t *testing.T) {
	tree1 := NewBinaryTree()

	if !tree1.Insert(TreeNode(3)) {
		t.Fatal()
	}
	if !tree1.Insert(TreeNode(-1)) {
		t.Fatal()
	}

	if !tree1.Insert(TreeNode(2)) {
		t.Fatal(tree1.String())
	}

	if !tree1.Insert(TreeNode(5)) {
		t.Fatal(tree1.String())
	}

	node := tree1.find(TreeNode(-1))
	if node == nil {
		t.Fatal()
	}

	max := tree1.Max()
	if max != TreeNode(5) {
		t.Fatal()
	}
	min := tree1.Min()
	if min != TreeNode(-1) {
		t.Fatal()
	}

	tree1.Delete(TreeNode(5))
	t.Log(tree1)
}
