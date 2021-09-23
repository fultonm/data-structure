package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := new(list.List)
	l.PushFront(1.2)
	myString := l.Front().Value
	fmt.Printf("%v", myString)
}
