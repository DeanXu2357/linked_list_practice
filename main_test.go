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

func TestRemoveListNode(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode(&l, &b)

	if l.ToString() != "3 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 1", l.ToString())
	}
}

func TestRemoveListNode_NodeNotExist(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	d := Node{next: nil, value: 4}

	RemoveListNode(&l, &d)

	if l.ToString() != "3 -> 2 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 2 -> 1", l.ToString())
	}
}

func TestRemoveListNode_Head(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode(&l, &c)

	if l.ToString() != "2 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "2 -> 1", l.ToString())
	}
}

func TestRemoveListNode_Last(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode(&l, &a)

	if l.ToString() != "3 -> 2" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 2", l.ToString())
	}
}

func TestRemoveListNode_OnlyOne(t *testing.T) {
	a := Node{next: nil, value: 1}

	l := List{head: &a}

	RemoveListNode(&l, &a)

	if l.ToString() != "" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}

func TestRemoveListNode_RemoveEmpty(t *testing.T) {
	l := List{head: nil}

	RemoveListNode(&l, nil)

	if l.ToString() != "" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}

func TestRemoveListNode2(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode2(&l, &b)

	if l.ToString() != "3 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 1", l.ToString())
	}
}

func TestRemoveListNode2_NodeNotExist(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	d := Node{next: nil, value: 4}

	RemoveListNode2(&l, &d)

	if l.ToString() != "3 -> 2 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 2 -> 1", l.ToString())
	}
}

func TestRemoveListNode2_Head(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode2(&l, &c)

	if l.ToString() != "2 -> 1" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "2 -> 1", l.ToString())
	}
}

func TestRemoveListNode2_Last(t *testing.T) {
	a := Node{next: nil, value: 1}
	b := Node{next: &a, value: 2}
	c := Node{next: &b, value: 3}

	l := List{head: &c}

	RemoveListNode2(&l, &a)

	if l.ToString() != "3 -> 2" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "3 -> 2", l.ToString())
	}
}

func TestRemoveListNode2_OnlyOne(t *testing.T) {
	a := Node{next: nil, value: 1}

	l := List{head: &a}

	RemoveListNode2(&l, &a)

	if l.ToString() != "" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}

func TestRemoveListNode2_RemoveEmpty(t *testing.T) {
	l := List{head: nil}

	RemoveListNode2(&l, nil)

	if l.ToString() != "" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}

func TestRemoveListNode2_ListIsNil(t *testing.T) {
	a := Node{next: nil, value: 1}
	var l *List
	RemoveListNode2(l, &a)

	if l != nil {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}

func TestRemoveListNode2_HeadIsNil(t *testing.T) {
	a := Node{next: nil, value: 1}

	l := List{head: nil}

	RemoveListNode2(&l, &a)

	if l.ToString() != "" {
		t.Errorf("List ToString() error: expect '%s', got '%s'", "", l.ToString())
	}
}
