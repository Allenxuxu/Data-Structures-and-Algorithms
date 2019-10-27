package hashtable

import "hash/fnv"

var loadFactor = 1.0

type item struct {
	key  string
	data interface{}
	next *item
}

// HashTable 取余法（定值）+拉链法（解决冲突）
// 未处理相同 key
type HashTable struct {
	data  []*item
	size  uint32
	count uint32
}

func New(initSize int) *HashTable {
	return &HashTable{
		data: make([]*item, initSize),
		size: uint32(initSize),
	}
}

func (t *HashTable) hash(s string) uint32 {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(s))
	return hash.Sum32() % t.size
}

func (t *HashTable) resize() {
	tmpHashTable := New(int(t.size * 2))

	for i := 0; i < len(t.data); i++ {
		current := t.data[i]
		for current != nil {
			tmpHashTable.Add(current.key, current.data)

			current = current.next
		}
	}

	t.data = tmpHashTable.data
	t.size = tmpHashTable.size
	t.count = tmpHashTable.count
}

func (t *HashTable) Add(key string, value interface{}) {
	if float64(t.count+1)/float64(t.size) > loadFactor {
		t.resize()
	}

	position := t.hash(key)
	current := t.data[position]
	if current == nil {
		t.data[position] = &item{key: key, data: value}
	} else {
		for current.next != nil {
			current = current.next
		}
		current.next = &item{key: key, data: value}
	}
	t.count++
}

func (t *HashTable) Get(key string) (interface{}, bool) {
	position := t.hash(key)
	current := t.data[position]
	for current != nil {
		if current.key == key {
			return current.data, true
		}
		current = current.next
	}
	return nil, false
}

func (t *HashTable) Set(key string, value interface{}) bool {
	position := t.hash(key)
	current := t.data[position]
	for current != nil {
		if current.key == key {
			current.data = value
			return true
		}
		current = current.next
	}
	return false
}

func (t *HashTable) Del(key string) bool {
	position := t.hash(key)
	current := t.data[position]
	if current != nil {
		if current.key == key {
			t.data[position] = current.next
			t.count--
			return true
		} else {
			for current.next != nil {
				if current.next.key == key {
					current.next = current.next.next
					t.count--
					return true
				}
				current = current.next
			}
		}
	}
	return false
}

func (t *HashTable) Length() int {
	return int(t.count)
}
