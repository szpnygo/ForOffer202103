package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var newList *ListNode
	p := head
	for p != nil {
		temp := p.Next
		p.Next = newList
		newList = p
		p = temp
	}
	return newList
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	newHead := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return newHead
// }

func removeElements(head *ListNode, val int) *ListNode {
	p := &ListNode{Next: head}
	r := p
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}
	return r.Next
}

func oddEvenList(head *ListNode) *ListNode {
	p := head
	r := &ListNode{}
	rs := r
	for p != nil {
		r.Next = p.Next
		r = r.Next
		if p.Next != nil {
			p.Next = p.Next.Next
		} else {
			p.Next = nil
		}
		if p.Next == nil {
			p.Next = rs.Next
			break
		}
		p = p.Next
	}

	return head
}

func isPalindrome(head *ListNode) bool {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
}
