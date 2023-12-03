package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	curr := result
	remainder := 0
	for {
		if l1 != nil {
			remainder += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remainder += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: remainder % 10, Next: nil}
		remainder = remainder / 10
		curr = curr.Next

		if l1 == nil && l2 == nil {
			if remainder != 0 {

				curr.Next = &ListNode{Val: remainder, Next: nil}
			}
			break
		}

	}
	return result.Next
}
