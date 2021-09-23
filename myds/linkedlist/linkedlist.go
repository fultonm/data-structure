package linkedlist

type LinkedList[T comparable] struct {
	prev *LinkedList[T]
	next *LinkedList[T]
	data T
}

func Add[T comparable](l LinkedList[T], ele T) bool {
	return true
}
