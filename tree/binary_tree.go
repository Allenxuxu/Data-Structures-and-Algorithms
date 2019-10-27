package tree

type Node struct {
	data   interface{}
	parent *Node
	left   *Node
	right  *Node
}

type BTree struct {
	root *Node
}

func NewBTree() *BTree {
	return &BTree{}
}

func (t *BTree) Inset(data interface{}) {
	node := &Node{data: data}
	if t.root == nil {
		t.root = node
		node.parent = t.root
	} else {

	}
}
