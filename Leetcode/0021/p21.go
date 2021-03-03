package main

import (
	"fmt"
	"strconv"
)

func main() {
	l1 := &ListNode{Val: 1}
	l1.appendListNode(2)
	l1.appendListNode(4)

	l2 := &ListNode{Val: 1}
	l2.appendListNode(3)
	l2.appendListNode(4)
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{Next: nil}
	head := list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next = l1
			l1 = l1.Next
		} else {
			list.Next = l2
			l2 = l2.Next
		}
		list = list.Next
	}
	if l1 != nil {
		list.Next = l1
	} else {
		list.Next = l2
	}
	return head.Next
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(" " + strconv.Itoa(list.Val))
		list = list.Next
	}
	fmt.Println()
}

func (list *ListNode) appendListNode(val int) {
	temp := list
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val}
}
