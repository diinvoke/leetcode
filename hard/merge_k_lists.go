package hard

/*
https://leetcode.com/problems/merge-k-sorted-lists/

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	size := len(lists)
	interval := 1
	for interval < size {
		for i := 0; i < size-interval; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
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
