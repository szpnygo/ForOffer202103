package main

import (
	"fmt"
	"strconv"
)

func main() {

	l1 := &ListNode{Val: 9}
	l1.appendListNode(9)
	l1.appendListNode(9)
	l1.appendListNode(9)
	l1.appendListNode(9)
	l1.appendListNode(9)
	l1.appendListNode(9)

	l2 := &ListNode{Val: 9}
	l2.appendListNode(9)
	l2.appendListNode(9)
	l2.appendListNode(9)

	printList(addTwoNumbers(l1, l2))

}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headL := l1
	result := l1
	for l1 != nil && l2 != nil {
		l1.Val += l2.Val
		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for headL != nil {
		if headL.Val >= 10 {
			headL.Val -= 10
			if headL.Next != nil {
				headL.Next.Val++
			} else {
				headL.Next = &ListNode{Val: 1}
			}
		}
		headL = headL.Next
	}
	return result
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
