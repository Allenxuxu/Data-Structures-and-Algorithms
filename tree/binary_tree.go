package tree

import (
	"strings"

	"github.com/Allenxuxu/dsa/queue"
)

type Element interface {
	Less(e Element) bool
	String() string
}

type Node struct {
	data   Element
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(data Element) bool {
	node := &Node{data: data}
	if t.root == nil {
		t.root = node
	} else {
		for current := t.root; current != nil; {
			if current.data == node.data {
				return false
			}

			if node.data.Less(current.data) {
				if current.left == nil {
					current.left = node
					node.parent = current
					break
				} else {
					current = current.left
				}
			} else {
				if current.right == nil {
					current.right = node
					node.parent = current
					break
				} else {
					current = current.right
				}
			}
		}
	}
	return true
}

func (t *BinaryTree) Delete(data Element) {
	toDelete := t.find(data)
	if toDelete == nil {
		return
	}

	if toDelete.left != nil && toDelete.right != nil {
		next := min(toDelete.right)
		if next != toDelete.right {
			next.parent.left = next.right
			if next.right != nil {
				next.right.parent = next.parent
			}

			next.right = toDelete.right
			toDelete.right.parent = next
			toDelete.parent.left = next
			next.left = toDelete.left
			next.parent = toDelete.parent
		} else {
			toDelete.parent.left = next
			next.parent = toDelete.parent
			next.left = toDelete.left
			toDelete.left.parent = next
		}
	} else if toDelete.left != nil {
		if toDelete.parent.left == toDelete {
			toDelete.parent.left = toDelete.left
			toDelete.left.parent = toDelete.parent
		} else if toDelete.parent.right == toDelete {
			toDelete.parent.right = toDelete.left
			toDelete.left.parent = toDelete.parent
		}
	} else if toDelete.right != nil {
		if toDelete.parent.left == toDelete {
			toDelete.parent.left = toDelete.right
			toDelete.left.right = toDelete.parent
		} else if toDelete.parent.right == toDelete {
			toDelete.parent.right = toDelete.right
			toDelete.left.right = toDelete.parent
		}
	} else {
		if toDelete.parent.left == toDelete {
			toDelete.parent.left = nil
		} else if toDelete.parent.right == toDelete {
			toDelete.parent.right = nil
		}
	}

	toDelete.parent = nil
	toDelete.left = nil
	toDelete.right = nil
}

func min(root *Node) *Node {
	if root != nil {
		for current := root; current != nil; {
			if current.left == nil {
				return current
			} else {
				current = current.left
			}
		}
	}

	return nil
}
func (t *BinaryTree) Min() Element {
	node := min(t.root)
	if node != nil {
		return node.data
	}
	return nil
}

func max(root *Node) *Node {
	if root != nil {
		for current := root; current != nil; {
			if current.right == nil {
				return current
			} else {
				current = current.right
			}
		}
	}

	return nil
}

func (t *BinaryTree) Max() Element {
	node := max(t.root)
	if node != nil {
		return node.data
	}
	return nil
}

func (t *BinaryTree) find(data Element) *Node {
	q := queue.LinkQueue{}
	// 层次遍历（广度优先）
	for current := t.root; current != nil; {
		if current.data == data {
			return current
		} else {
			if current.left != nil {
				q.Push(current.left)
			}
			if current.right != nil {
				q.Push(current.right)
			}
		}
		e, ok := q.Pop()
		if ok {
			current = e.(*Node)
		} else {
			current = nil
		}
	}
	return nil
}

func (t *BinaryTree) String() string {
	var ret []string
	q := queue.LinkQueue{}
	// 层次遍历（广度优先）
	for current := t.root; current != nil; {
		ret = append(ret, current.data.String())

		if current.left != nil {
			q.Push(current.left)
		}
		if current.right != nil {
			q.Push(current.right)
		}

		e, ok := q.Pop()
		if ok {
			current = e.(*Node)
		} else {
			current = nil
		}
	}

	return strings.Join(ret, ",")
}

func (t *BinaryTree) PreOrderTraversal(node *Node, q *queue.LinkQueue) {
	if node != nil {
		q.Push(node.data)
		t.PreOrderTraversal(node.left, q)
		t.PreOrderTraversal(node.right, q)
	}
}

func (t *BinaryTree) InOrderTraversal(node *Node, q *queue.LinkQueue) {
	if node != nil {
		t.InOrderTraversal(node.left, q)
		q.Push(node.data)
		t.InOrderTraversal(node.right, q)
	}
}

func (t *BinaryTree) PostOrderTraversal(node *Node, q *queue.LinkQueue) {
	if node != nil {
		t.PostOrderTraversal(node.left, q)
		t.PostOrderTraversal(node.right, q)
		q.Push(node.data)
	}
}
