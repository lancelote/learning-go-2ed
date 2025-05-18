package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Next  *Node[T]
	Value T
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func (l *List[T]) Insert(t T, pos int) {
	dummy := &Node[T]{
		Next: l.Head,
	}

	cur := dummy
	i := pos

	for cur != nil && i != 0 {
		cur = cur.Next
		i--
	}

	if cur == nil || cur.Next == nil {
		l.Add(t)
		return
	}

	n := &Node[T]{
		Value: t,
	}

	tmp := cur.Next
	cur.Next = n
	n.Next = tmp
	l.Head = dummy.Next
}

func PrintEls[T comparable](l *List[T]) {
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}

func main() {
	// 100 200 5 10

	l := &List[int]{}
	l.Add(5)
	l.Add(10)

	fmt.Println(l.Index(5))  // 0
	fmt.Println(l.Index(10)) // 1
	fmt.Println(l.Index(20)) // -1

	l.Insert(100, 0)

	fmt.Println(l.Index(5))   // 1
	fmt.Println(l.Index(10))  // 2
	fmt.Println(l.Index(20))  // -1
	fmt.Println(l.Index(100)) // 0

	l.Insert(200, 1)

	fmt.Println(l.Index(5))   // 2
	fmt.Println(l.Index(10))  // 3
	fmt.Println(l.Index(200)) // 1
	fmt.Println(l.Index(20))  // -1
	fmt.Println(l.Index(100)) // 0

	PrintEls(l)
	l.Insert(300, 10)
	PrintEls(l)
	l.Add(400)
	PrintEls(l)
	l.Insert(500, 6)
	PrintEls(l)
}
