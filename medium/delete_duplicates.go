package medium

/*
82. Remove Duplicates from Sorted List II https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Return the linked list sorted as well.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
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

	pre := new(ListNode)
	newHead := pre
	i := head
	skip := 0

	for i != nil {
		if i.Next != nil && i.Val == i.Next.Val {
			skip += 1
			i = i.Next
			continue
		}
		if skip == 0 {
			pre.Next = i
			pre = pre.Next
			i = i.Next
			pre.Next = nil
			continue
		}
		if i.Next != nil && ((i.Next.Next != nil && i.Next.Val != i.Next.Next.Val) || i.Next.Next == nil) {
			pre.Next = i.Next
			i = i.Next.Next
			pre = pre.Next
			pre.Next = nil
			skip = 0
			continue
		}
		i = i.Next
	}

	return newHead.Next
}
