package medium

/*
92. Reverse Linked List II https://leetcode.com/problems/reverse-linked-list-ii/

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return dummy.Next
}
