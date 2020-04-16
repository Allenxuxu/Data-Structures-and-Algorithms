package stack

import "github.com/Allenxuxu/dsa/list"

// LinkStack 链表实现的堆栈
type LinkStack struct {
	space list.LinkList
}

func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func (s *LinkStack) Push(i interface{}) {
	s.space.PushBack(i)
}

func (s *LinkStack) Pop() (interface{}, bool) {
	return s.space.PopBack()
}

func (s *LinkStack) Peek() (interface{}, bool) {
	return s.space.Get(s.space.Length() - 1)
}

func (s *LinkStack) Clear() {
	s.space.Clear()
}

func (s *LinkStack) Length() int {
	return s.space.Length()
}
