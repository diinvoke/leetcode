package medium

/*
86. Partition List https://leetcode.com/problems/partition-list/

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
*/

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	lessHead := new(ListNode)
	lessDummy := lessHead

	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val < x {
			lessDummy.Next = cur.Next
			lessDummy = lessDummy.Next
			cur.Next = cur.Next.Next
			lessDummy.Next = nil
		} else {
			cur = cur.Next
		}
	}
	lessDummy.Next = dummy.Next
	return lessHead.Next
}
