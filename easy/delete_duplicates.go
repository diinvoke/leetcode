package easy

/*
83. Remove Duplicates from Sorted List https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l, r := head, head

	for r != nil {
		if l.Val == r.Val {
			r = r.Next
			if r == nil {
				l.Next = nil
			}
		} else {
			l.Next = r
			l = r
			r = l
		}
	}

	return head
}
