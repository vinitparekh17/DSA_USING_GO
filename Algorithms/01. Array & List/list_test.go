package main

import (
	"container/list"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := list.New()
	linkedList.PushFront(1)
	if linkedList.Front().Value != 1 {
		t.Error("PushFront() failed")
	}
	linkedList.PushBack(2)
	if linkedList.Back().Value != 2 {
		t.Error("PushBack() failed")
	}
	if linkedList.Front().Next().Value != 2 {
		t.Error("Next() failed")
	}
	anotherList := list.New()
	anotherList.PushBack(3)
	anotherList.PushBack(4)
	linkedList.PushBackList(anotherList)
	if linkedList.Back().Value != 4 {
		t.Error("PushBackList() failed")
	}
	if linkedList.Front().Next().Next().Value != 3 {
		t.Error("PushBackList() failed")
	}
	linkedList.Remove(linkedList.Front().Next().Next())
	if linkedList.Front().Next().Next().Value != 4 {
		t.Error("Remove() failed")
	}
}
