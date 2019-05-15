package easy

import (
	"fmt"
	"strings"
)

func genListNode(nums []int) *ListNode {
	head := new(ListNode)
	current := head
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head.Next
}

func genCycleList(nums []int, pos int) *ListNode {
	list := genListNode(nums)
	if pos == -1 {
		return list
	}

	head := list
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	tail := arr[len(arr)-1]
	tail.Next = arr[pos]

	return list
}

func stringListNode(l *ListNode) string {
	var nums []string
	for l != nil {
		nums = append(nums, fmt.Sprintf("%d", l.Val))
		l = l.Next
	}

	return strings.Join(nums, ",")
}

func equalListNode(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}
