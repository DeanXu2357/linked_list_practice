package main

import "testing"

func Test_List_ToString(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	if l.ToString() != "3 -> 2 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 2 -> 1", l.ToString())
	}
}
