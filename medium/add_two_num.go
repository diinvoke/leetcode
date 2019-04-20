package medium

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var carry int

	currentNode := &head

	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		sum := l1Val + l2Val + carry
		carry = sum / 10
		currentNode.Val = sum % 10

		if l1 != nil || l2 != nil {
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next
		}
	}

	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}

	return &head
}
