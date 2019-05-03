package medium

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Given a linked list, remove the n-th node from the end of list and return its head.

Example:
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.

Follow up:
Could you do this in one pass?
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	quick, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		quick = quick.Next
	}

	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
