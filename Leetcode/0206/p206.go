package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	newList := *ListNode
	p := head
	for p != nil {
		temp := p.Next
		p.Next = newList
		newList = p
		p = temp
	}
	return newList.Next
}
