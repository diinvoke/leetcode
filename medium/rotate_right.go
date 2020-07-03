package medium

/*
61. Rotate List https://leetcode.com/problems/rotate-list/
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL

*/

func rotateRight(head *ListNode, k int) *ListNode {
	head = reverseList(head)
	var tail *ListNode
	for p := head; p != nil; p = p.Next {
		tail = p
	}
	if tail == nil || head.Next == nil {
		return head
	}

	for i := 0; i < k; i++ {
		headNext := head.Next
		head.Next = nil
		tail.Next = head
		tail = head
		head = headNext
	}

	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	j := p.Next
	p.Next = nil

	for {
		jNext := j.Next
		j.Next = p
		p = j
		if jNext == nil {
			break
		}
		j = jNext
	}

	return j
}

func rotateRight3(head *ListNode, k int) *ListNode {
	var tail *ListNode
	total := 0
	for p := head; p != nil; p = p.Next {
		total += 1
		tail = p
	}

	if tail == nil || head.Next == nil {
		return head
	}

	k %= total
	f, s := head, head
	for i := 0; i < k; i++ {
		f = f.Next
	}

	for f.Next != nil {
		f = f.Next
		s = s.Next
	}

	f.Next = head
	f = s.Next
	s.Next = nil
	return f
}

func rotateRight4(head *ListNode, k int) *ListNode {
	var tail *ListNode
	total := 0
	for p := head; p != nil; p = p.Next {
		total += 1
		tail = p
	}

	if tail == nil || head.Next == nil {
		return head
	}

	tail.Next = head
	loss := total - k%total
	for i := 0; i < loss; i++ {
		tail = tail.Next
	}

	head = tail.Next
	tail.Next = nil
	return head
}
