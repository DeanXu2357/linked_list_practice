package main

import (
	"fmt"
)

func main() {
	a := Node{value: "a"}
	b := Node{next: &a, value: "b"}
	c := Node{next: &b, value: "c"}
	l := List{head: &c}

	fmt.Println(l.ToString())
}

type Node struct {
	next  *Node
	value any
}

type List struct {
	head *Node
}

func (l *List) ToString() string {
	var output string
	list := l.head
	for list != nil {
		output = fmt.Sprintf("%s%v", output, list.value)
		if list.next != nil {
			output = fmt.Sprintf("%s -> ", output)
		}
		list = list.next
	}
	return output
}

//// 最一開始 copilot 給我的
//func RemoveListNode(l *List, n *Node) {
//	if n == l.head {
//		l.head = n.next
//	} else {
//		list := l.head
//		for list.next != n && list.next != nil {
//			list = list.next
//		}
//		list.next = n.next
//	}
//}

// RemoveListNode removes a node from a list
// 註：copilot 協助完成
func RemoveListNode(l *List, n *Node) {
	if l == nil || n == nil || l.head == nil {
		return
	}

	if n == l.head {
		l.head = n.next
		return
	}

	list := l.head
	for list.next != n && list.next != nil {
		list = list.next
	}

	if list.next == n {
		list.next = n.next
		return
	}
}

// RemoveListNode2 removes a node from a list
// 註：自己寫的 Linus taste version
func RemoveListNode2(l *List, n *Node) {
	var current **Node = &l.head
	for *current != n && *current != nil {
		current = &(*current).next
	}

	*current = n.next
}
