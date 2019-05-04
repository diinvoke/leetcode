package easy

/*
https://leetcode.com/problems/merge-sorted-array/

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

Example:
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

func MergeArray(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	storageIndex := m + n - 1
	m -= 1
	n -= 1

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[storageIndex] = nums1[m]
			m -= 1
		} else {
			nums1[storageIndex] = nums2[n]
			n -= 1
		}

		storageIndex -= 1
	}

	for m >= 0 {
		nums1[storageIndex] = nums1[m]
		m -= 1
		storageIndex -= 1
	}

	for n >= 0 {
		nums1[storageIndex] = nums2[n]
		n -= 1
		storageIndex -= 1
	}
}
