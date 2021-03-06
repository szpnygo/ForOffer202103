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
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
//     seen := map[*ListNode]struct{}{}
//     for head != nil {
//         if _, ok := seen[head]; ok {
//             return head
//         }
//         seen[head] = struct{}{}
//         head = head.Next
//     }
//     return nil
// }

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	seen := map[*ListNode]struct{}{}
// 	for headA != nil {
// 		seen[headA] = struct{}{}
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if _, ok := seen[headB]; ok {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}
// 	return nil
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		a = a.Next
		b = b.Next
		if a == nil && b == nil {
			return a
		}
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
	}
	return a
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{Next: head}
	slow, fast := h, h
	for fast != nil {
		if n < 0 {
			slow = slow.Next
		} else {
			n--
		}
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return h.Next
}
