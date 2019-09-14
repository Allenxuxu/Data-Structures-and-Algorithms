package stack

// StaticStack 数组实现的堆栈
type StaticStack struct {
	top   int
	space []interface{}
}

func NewStaticStack(size int) *StaticStack {
	return &StaticStack{
		space: make([]interface{}, size),
	}
}

func (s *StaticStack) Push(i interface{}) bool {
	if s.top == len(s.space) {
		return false
	}
	s.space[s.top] = i
	s.top++

	return true
}

func (s *StaticStack) Pop() (interface{}, bool) {
	if s.top == 0 {
		return nil, false
	}
	s.top--
	i := s.space[s.top]

	return i, true
}

func (s *StaticStack) Peek() (interface{}, bool) {
	if s.top == 0 {
		return nil, false
	}
	return s.space[s.top-1], true
}

func (s *StaticStack) Clear() {
	s.top = 0
}

func (s *StaticStack) Capacity() int {
	return len(s.space)
}

func (s *StaticStack) Length() int {
	return s.top
}
