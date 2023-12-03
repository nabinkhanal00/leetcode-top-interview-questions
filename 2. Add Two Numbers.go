package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	initial := l1
	var prev *ListNode = nil
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}
		l1.Val = sum
		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}

		l1.Val = sum
		prev = l1
		l1 = l1.Next
	}
	for l2 != nil {

		sum := l1.Val + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}
		prev.Next = &ListNode{Val: sum, Next: nil}
		prev = prev.Next
		l2 = l2.Next
	}
	if carry != 0 {
		prev.Next = &ListNode{Val: carry, Next: nil}
	}
	return initial
}
