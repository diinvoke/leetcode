package easy

/*
https://leetcode.com/problems/merge-two-sorted-lists/

Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	dummy := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			dummy.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}

		dummy = dummy.Next
	}

	for l1 != nil {
		dummy.Next = &ListNode{Val: l1.Val}
		dummy = dummy.Next
		l1 = l1.Next
	}

	for l2 != nil {
		dummy.Next = &ListNode{Val: l2.Val}
		dummy = dummy.Next
		l2 = l2.Next
	}

	return head.Next
}
