package tree

import (
	"strconv"
	"testing"

	"github.com/Allenxuxu/dsa/queue"
	"github.com/stretchr/testify/assert"
)

type TreeNode int

func (t TreeNode) Less(e Element) bool {
	return t < e.(TreeNode)
}

func (t TreeNode) String() string {
	return strconv.Itoa(int(t))
}

func TestBinaryTree_Insert(t *testing.T) {
	tree := NewBinaryTree()

	if !tree.Insert(TreeNode(3)) {
		t.Fatal()
	}
	if !tree.Insert(TreeNode(-1)) {
		t.Fatal()
	}
	if !tree.Insert(TreeNode(2)) {
		t.Fatal(tree.String())
	}
	if !tree.Insert(TreeNode(5)) {
		t.Fatal(tree.String())
	}
	//       3
	//  -1      5
	//     2
	q := queue.NewLinkQueue()
	tree.LevelTraversal(tree.root, q)
	length := q.Length()
	var result []int
	for i := 0; i < length; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatal()
		}
		result = append(result, int(v.(TreeNode)))
	}

	assert.Equal(t, []int{3, -1, 5, 2}, result)

	node := tree.find(TreeNode(-1))
	if node == nil {
		t.Fatal()
	}

	max := tree.Max()
	if max != TreeNode(5) {
		t.Fatal()
	}
	min := tree.Min()
	if min != TreeNode(-1) {
		t.Fatal()
	}

	tree.Delete(TreeNode(5))
	t.Log(tree)
}

func TestBinaryTree_PreOrderTraversal(t *testing.T) {
	tree := NewBinaryTree()

	data := []int{5, 1, 4, 3, 6}
	//       5
	//  1       6
	//     4
	//  3
	for i := 0; i < len(data); i++ {
		tree.Insert(TreeNode(data[i]))
	}

	q := queue.NewLinkQueue()

	tree.PreOrderTraversal(tree.root, q)
	length := q.Length()
	var result []int
	for i := 0; i < length; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatal()
		}
		result = append(result, int(v.(TreeNode)))
	}

	assert.Equal(t, []int{5, 1, 4, 3, 6}, result)
}

func TestBinaryTree_InOrderTraversal(t *testing.T) {
	tree := NewBinaryTree()

	data := []int{5, 1, 4, 3, 6}
	//       5
	//  1       6
	//     4
	//  3
	for i := 0; i < len(data); i++ {
		tree.Insert(TreeNode(data[i]))
	}

	q := queue.NewLinkQueue()

	tree.InOrderTraversal(tree.root, q)
	length := q.Length()
	var result []int
	for i := 0; i < length; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatal()
		}
		result = append(result, int(v.(TreeNode)))
	}

	assert.Equal(t, []int{1, 3, 4, 5, 6}, result)
}

func TestBinaryTree_PostOrderTraversal(t *testing.T) {
	tree := NewBinaryTree()

	data := []int{5, 1, 4, 3, 6}
	//       5
	//  1       6
	//     4
	//  3
	for i := 0; i < len(data); i++ {
		tree.Insert(TreeNode(data[i]))
	}

	q := queue.NewLinkQueue()

	tree.PostOrderTraversal(tree.root, q)
	length := q.Length()
	var result []int
	for i := 0; i < length; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatal()
		}
		result = append(result, int(v.(TreeNode)))
	}

	assert.Equal(t, []int{3, 4, 1, 6, 5}, result)
}

func TestBinaryTree_LevelTraversal(t *testing.T) {
	tree := NewBinaryTree()

	data := []int{5, 1, 4, 3, 6}
	//       5
	//  1       6
	//     4
	//  3
	for i := 0; i < len(data); i++ {
		tree.Insert(TreeNode(data[i]))
	}

	q := queue.NewLinkQueue()
	tree.LevelTraversal(tree.root, q)
	length := q.Length()
	var result []int
	for i := 0; i < length; i++ {
		v, ok := q.Pop()
		if !ok {
			t.Fatal()
		}
		result = append(result, int(v.(TreeNode)))
	}

	assert.Equal(t, []int{5, 1, 6, 4, 3}, result)
}
