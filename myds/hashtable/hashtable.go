package hashtable

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math"
)

func HashFunc(key int, colSize int) int {
	return key % colSize
}

func LinearProbe(i int) int {
	return i
}

func QuadraticProbe(i int) int {
	return pow(i, 2)
}

type Node[T comparable] struct {
	key   int
	value T
}

type HashingFunc func(int, int) int

type ProbingFunc func(int) int

type HashTable[T comparable] struct {
	col       []*Node[T]
	colSize   int
	len       int
	hashFunc  HashingFunc
	probeFunc ProbingFunc
}

func (ht *HashTable[T]) Init(colSize int, hashFunc HashingFunc, probeFunc ProbingFunc) *HashTable[T] {
	ht.col = make([]*Node[T], colSize)
	ht.colSize = colSize
	ht.hashFunc = hashFunc
	ht.probeFunc = probeFunc
	return ht
}

func New[T comparable](length int, hashFunc HashingFunc, probeFunc ProbingFunc) *HashTable[T] {
	return new(HashTable[T]).Init(length, hashFunc, probeFunc)
}

func (ht *HashTable[T]) insertIntKey(key int, value T) (*HashTable[T], error) {
	if float64(ht.len) > float64(ht.colSize)*0.8 {
		return ht, errors.New(fmt.Sprintf("Hashtable over capacity. Resizing not implemented."))
	}
	i := 0
	for {
		colIdx := getColIndex(ht, key, i)
		if ht.col[colIdx] == nil {
			ht.col[colIdx] = &Node[T]{key, value}
			ht.len++
			break
		}
		i++

	}
	return ht, nil
}

func (ht *HashTable[T]) Insert(key string, value T) (*HashTable[T], error) {
	return ht.insertIntKey(hash(key), value)
}

func (ht *HashTable[T]) getByKey(key int) *T {
	node := ht.col[getColIndex(ht, key, 0)]
	i := 1
	for node != nil && node.key != key {
		node = ht.col[getColIndex(ht, key, 0)]
		i++
	}
	if node == nil {
		return nil
	}
	return &node.value
}

func (ht *HashTable[T]) Get(key string) *T {
	return ht.getByKey(hash(key))
}

func (ht *HashTable[T]) Len() int {
	return ht.len
}

func getColIndex[T comparable](ht *HashTable[T], key int, i int) int {
	return (ht.hashFunc(key, ht.colSize) + ht.probeFunc(i)) % ht.colSize
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

// https://stackoverflow.com/questions/13582519
// FNV32a hashes using fnv32a algorithm
func hash(text string) int {
	algorithm := fnv.New32()
	algorithm.Write([]byte(text))
	return int(math.Abs((float64(algorithm.Sum32()))))
}
