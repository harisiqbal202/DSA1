package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedlist) printListData() {
	toprint := l.head
	for l.length != 0 {
		fmt.Println(toprint.data)
		toprint = toprint.next
		l.length--
	}
	fmt.Println("")
}

func (l *linkedlist) deletevalue(val int) {
	if l.length == 0 {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousHeadToDelete := l.head
	for previousHeadToDelete.next.data != val {
		if previousHeadToDelete.next.next == nil {
			return
		}
		previousHeadToDelete = previousHeadToDelete.next
	}
	previousHeadToDelete = previousHeadToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 82}
	node3 := &node{data: 41}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)

	fmt.Println(mylist)
	mylist.printListData()
}
