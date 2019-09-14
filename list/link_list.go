package list

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type LinkList struct {
	head   Node
	length int
}

func (l *LinkList) PushBack(v interface{}) {
	node := &l.head
	for node.next != nil {
		node = node.next
	}

	data := &Node{data: v}
	if node != &l.head {
		data.prev = node
	}
	node.next = data
	l.length++
}

func (l *LinkList) PushFront(v interface{}) {
	data := &Node{data: v}
	data.next = l.head.next
	if l.head.next != nil {
		l.head.next.prev = data
	}
	l.head.next = data

	l.length++
}

func (l *LinkList) Insert(i int, v interface{}) bool {
	if i > l.length || i < 0 {
		return false
	}

	node := &l.head
	for k := 0; k < i; k++ {
		node = node.next
	}
	data := &Node{data: v}
	if i != 0 {
		data.prev = node
	}
	data.next = node.next
	node.next.prev = data
	node.next = data
	l.length++

	return true
}

func (l *LinkList) Get(i int) (interface{}, bool) {
	if i > l.length || i < 0 {
		return nil, false
	}

	node := l.head.next
	for k := 0; k < i; k++ {
		node = node.next
	}

	return node.data, true
}

func (l *LinkList) PopBack() (interface{}, bool) {
	if l.length == 0 {
		return nil, false
	}

	node := l.head.next
	for node.next != nil {
		node = node.next
	}
	if node.prev == nil {
		l.head.next = nil
	} else {
		node.prev.next = nil
		node.prev = nil
	}

	l.length--

	return node.data, true
}

func (l *LinkList) PopFront() (interface{}, bool) {
	if l.length == 0 {
		return nil, false
	}

	node := l.head.next
	if node.next != nil {
		node.next.prev = nil
	}
	l.head.next = node.next
	l.length--

	return node.data, true
}

func (l *LinkList) Find(v interface{}) (int, bool) {
	if l.length == 0 {
		return -1, false
	}

	node := l.head.next
	var index int
	for node != nil {
		if node.data == v {
			return index, true
		}
		node = node.next
		index++
	}

	return -1, false
}

func (l *LinkList) Length() int {
	return l.length
}
