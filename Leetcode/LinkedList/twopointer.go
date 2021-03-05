package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next.Next
	for slow.Next != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next.Next
	for slow.Next != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
