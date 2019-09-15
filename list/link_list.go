package list

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type LinkList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkList) PushBack(v interface{}) {
	data := &Node{data: v}

	if l.head == nil {
		l.head = data
		l.tail = data
	} else {
		l.tail.next = data
		data.prev = l.tail
		l.tail = data
	}

	l.length++
}

func (l *LinkList) PushFront(v interface{}) {
	data := &Node{data: v}

	if l.head == nil {
		l.head = data
		l.tail = data
	} else {
		l.head.prev = data
		data.next = l.head
		l.head = data
	}

	l.length++
}

func (l *LinkList) PopBack() (interface{}, bool) {
	if l.length == 0 {
		return nil, false
	}

	node := l.tail
	l.tail = node.prev
	if node.prev != nil {
		node.prev.next = nil
	}
	node.prev = nil

	if l.tail == nil {
		l.head = nil
	}

	l.length--
	return node.data, true
}

func (l *LinkList) PopFront() (interface{}, bool) {
	if l.length == 0 {
		return nil, false
	}

	node := l.head
	l.head = node.next
	if node.next != nil {
		node.next.prev = nil
	}
	node.next = nil

	if l.head == nil {
		l.tail = nil
	}

	l.length--
	return node.data, true
}

func (l *LinkList) Find(v interface{}) (int, bool) {
	if l.length == 0 {
		return -1, false
	}

	node := l.head
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

func (l *LinkList) Insert(i int, v interface{}) bool {
	if i > l.length || i < 0 {
		return false
	}

	if i == 0 {
		l.PushFront(v)
		return true
	}
	if i == l.length {
		l.PushBack(v)
		return true
	}

	data := &Node{data: v}
	var node *Node
	if i < (l.length / 2) {
		node = l.head
		for k := 0; k < i; k++ {
			node = node.next
		}
	} else {
		node = l.tail
		index := l.length - 1 - i
		for k := 0; k < index; k++ {
			node = node.prev
		}
	}
	node.prev.next = data
	data.prev = node.prev
	node.next.prev = data
	data.next = node.next

	l.length++
	return true
}

func (l *LinkList) Get(i int) (interface{}, bool) {
	if i >= l.length || i < 0 {
		return nil, false
	}

	var node *Node
	if i < (l.length / 2) {
		node = l.head
		for k := 0; k < i; k++ {
			node = node.next
		}
	} else {
		node = l.tail
		index := l.length - 1 - i
		for k := 0; k < index; k++ {
			node = node.prev
		}
	}

	return node.data, true
}

func (l *LinkList) Length() int {
	return l.length
}

func (l *LinkList) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}
