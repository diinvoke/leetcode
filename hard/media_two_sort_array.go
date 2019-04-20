package hard

/*
here are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0, len(nums1)+len(nums2))
	len1, len2 := len(nums1), len(nums2)
	var index1, index2 int

	for index1 < len1 && index2 < len2 {
		v1 := nums1[index1]
		v2 := nums2[index2]
		if v1 < v2 {
			nums = append(nums, v1)
			index1 += 1
			continue
		}

		nums = append(nums, v2)
		index2 += 1
	}

	if index1 < len1 {
		nums = append(nums, nums1[index1:]...)
	}

	if index2 < len2 {
		nums = append(nums, nums2[index2:]...)
	}

	lenT := len1 + len2
	if lenT%2 != 0 {
		return float64(nums[lenT/2])
	}

	return (float64(nums[(lenT-1)/2]) + float64(nums[(lenT-1)/2+1])) / 2
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	lenTotal := len1 + len2
	nums := make([]int, lenTotal)
	var index1, index2 int

	for i := 0; i < lenTotal; i++ {
		if index1 == len1 && index2 < len2 {
			nums[i] = nums2[index2]
			index2 += 1
			continue
		}

		if index2 == len2 && index1 < len1 {
			nums[i] = nums1[index1]
			index1 += 1
			continue
		}

		if nums1[index1] < nums2[index2] {
			nums[i] = nums1[index1]
			index1 += 1
			continue
		}

		nums[i] = nums2[index2]
		index2 += 1
	}

	if lenTotal%2 != 0 {
		return float64(nums[lenTotal/2])
	}

	return (float64(nums[(lenTotal-1)/2]) + float64(nums[(lenTotal-1)/2+1])) / 2
}

func FindMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n
	if total%2 == 0 {
		return (float64(findMinK(nums1, m, nums2, n, total/2)) + float64(findMinK(nums1, m, nums2, n, total/2+1))) / 2
	}

	return float64(findMinK(nums1, m, nums2, n, total/2+1))
}

// nums1 and nums2 is asc sorted array and must be ensure m < n
func findMinK(nums1 []int, m int, nums2 []int, n int, k int) int {
	if m > n {
		return findMinK(nums2, n, nums1, m, k)
	}
	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	pa := min(k/2, m)
	pb := k - pa
	if nums1[pa-1] == nums2[pb-1] {
		return nums1[pa-1]
	}
	if nums1[pa-1] < nums2[pb-1] {
		return findMinK(nums1[pa:], m-pa, nums2, n, k-pa)
	}

	return findMinK(nums1, m, nums2[pb:], n-pb, k-pb)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
