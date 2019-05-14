package hard

/*
https://leetcode.com/problems/reverse-nodes-in-k-group/

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	first := head
	tail := head
	for i := 1; tail != nil; i++ {
		if i%k == 0 {
			tailNext := tail.Next
			tail.Next = nil
			pre.Next = reverseList(first)

			pre = first
			first = tailNext
			tail = tailNext
			continue
		}

		if tail == nil {
			break
		}
		tail = tail.Next
	}

	pre.Next = first

	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	q := p.Next
	var pre *ListNode

	for q != nil {
		l := q.Next
		q.Next = p
		p.Next = pre

		pre = p
		p = q
		q = l
	}

	return p
}
