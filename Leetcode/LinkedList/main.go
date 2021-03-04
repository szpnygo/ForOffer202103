package main

import (
	"fmt"
)

type MyLinkedList struct {
	val  int
	next *MyLinkedList
	len  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.len {
		return -1
	}
	for index != 0 {
		index--
		this = this.next
	}
	return this.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.len == 0 {
		this.val = val
	} else {
		list := &MyLinkedList{val: this.val, next: this.next}
		this.val = val
		this.next = list
	}
	this.len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.len == 0 {
		this.AddAtHead(val)
		return
	}
	head := this
	for head.next != nil {
		head = head.next
	}
	head.next = &MyLinkedList{val: val}
	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.len {
		return
	}
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	head := this
	for head != nil {
		index--
		if index == 0 {
			next := head.next
			head.next = &MyLinkedList{val: val, next: next}
			this.len++
			return
		}
		head = head.next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		next := this.next
		if next == nil {
			this.len--
		} else {
			this.val = next.val
			this.next = next.next
			this.len--
		}
	}
	if index >= this.len {
		return
	}
	head := this
	for head != nil {
		index--
		if index == 0 {
			head.next = head.next.next
			this.len--
			return
		}
		head = head.next
	}
}

func (this *MyLinkedList) Print() {
	for this != nil {
		fmt.Print(this.val, ",")
		this = this.next
	}
	fmt.Println("")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()
	list.AddAtTail(1)
	fmt.Println("GET", list.Get(0))

}
