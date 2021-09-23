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
	col      []Node[T]
	hashFunc HashingFunc
}

func Add(ht HashTable[T])

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
