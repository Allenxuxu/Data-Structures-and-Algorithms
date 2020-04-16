package queue

import "github.com/Allenxuxu/dsa/list"

type LinkQueue struct {
	space list.LinkList
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{}
}

func (q *LinkQueue) Push(v interface{}) {
	q.space.PushBack(v)
}

func (q *LinkQueue) Pop() (interface{}, bool) {
	return q.space.PopFront()
}

func (q *LinkQueue) Clear() {
	q.space.Clear()
}

func (q *LinkQueue) Length() int {
	return q.space.Length()
}
