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
