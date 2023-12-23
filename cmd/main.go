package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(val int) {
	if l.Head == nil {
		l.Head = &Node{Val: val}
		return
	}

	l.Head = &Node{
		Val:  val,
		Next: l.Head,
	}
}

func (l *LinkedList) Pop() int {
	if l.Head == nil {
		return 0
	}

	previous := l.Head
	l.Head = l.Head.Next

	return previous.Val
}

func (l *LinkedList) Remove() {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
}

func main() {
	linkedList := LinkedList{}

	for i := 0; i < 10; i++ {
		linkedList.Add(i + 1)
	}

	for linkedList.Head != nil {
		fmt.Println(linkedList.Pop())
	}
}
