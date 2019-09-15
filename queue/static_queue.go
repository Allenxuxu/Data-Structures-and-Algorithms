package queue

type StaticQueue struct {
	read   int
	write  int
	length int
	size   int
	space  []interface{}
}

func NewStaticQueue(size int) *StaticQueue {
	return &StaticQueue{
		size:  size,
		space: make([]interface{}, size),
	}
}

func (q *StaticQueue) Push(v interface{}) bool {
	if q.length >= q.size {
		return false
	}

	q.space[q.write] = v
	q.write = (q.write + 1) % q.size
	q.length++

	return true
}

func (q *StaticQueue) Pop() (interface{}, bool) {
	if q.length <= 0 {
		return nil, false
	}

	data := q.space[q.read]
	q.read = (q.read + 1) % q.size
	q.length--

	return data, true
}

func (q *StaticQueue) Clear() {
	q.length = 0
	q.write = 0
	q.read = 0
}

func (q *StaticQueue) Capacity() int {
	return q.size
}

func (q *StaticQueue) Length() int {
	return q.length
}
