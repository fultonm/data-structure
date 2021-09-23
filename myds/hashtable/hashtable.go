package hashtable

func HashFunc(key int, len int) int {
	return key % len
}

func LinearProbe(hash int, i int) int {
	return hash + i
}

func QuadraticProbe(hash int, i int) int {
	return hash + pow(i, 2)
}

type Node[T comparable] struct {
	key   int
	value T
}

type HashingFunc func(int, int) int

type ProbingFunc func(int, int) int

type HashTable[T comparable] struct {
	col       []Node[T]
	length    int
	hashFunc  HashingFunc
	probeFunc ProbingFunc
}

func New[T comparable](length int) *HashTable[T] {
	h := HashTable[T]{
		col:       make([]Node[T], length),
		length:    length,
		hashFunc:  HashFunc,
		probeFunc: LinearProbe,
	}
	return &h
}

func (ht HashTable[T]) Add(key int, value T) T {
	hash := ht.hashFunc(key, ht.length)
	if ht.col[hash] == nil {

	}
}

func pow(a int, b int) int {
	sign := 1
	if b < 0 {
		sign = -1
		b *= sign
	}
	r := 1
	for i := 0; i < b; i++ {
		r *= a
	}
	return r * sign
}
