package medium

/*
https://leetcode.com/problems/swap-nodes-in-pairs/


Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

func SwapPairs(head *ListNode) *ListNode {
	return swapPairs(head)
}

func swapPairs(head *ListNode) *ListNode {
	first := &ListNode{Next: head}

	dummy := first
	one := dummy.Next
	var two *ListNode
	if one != nil {
		two = one.Next
	}

	for two != nil {
		dummy.Next = two

		temp := two.Next
		two.Next = one
		one.Next = temp
		dummy = one
		one = one.Next
		if one == nil {
			break
		}
		two = one.Next
	}

	return first.Next
}
