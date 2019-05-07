package tools

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}

	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	storeIndex := left

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex += 1
		}
	}

	nums[right], nums[storeIndex] = nums[storeIndex], nums[right]
	return storeIndex
}

func HeapSort(nums []int) []int {
	n := len(nums) - 1
	for k := n / 2; k >= 0; k-- {
		heapSink(nums, k, n)
	}

	for n >= 0 {
		nums[0], nums[n] = nums[n], nums[0]
		n -= 1
		heapSink(nums, 0, n)
	}

	return nums
}

func heapSink(nums []int, k, n int) {
	for 2*k+1 <= n {
		j := 2*k + 1
		if j < n && nums[j] < nums[j+1] {
			j += 1
		}
		if nums[k] >= nums[j] {
			break
		}
		nums[k], nums[j] = nums[j], nums[k]
		k = j
	}
}

func heapSwim(nums []int, k int) {
	for k > 0 && nums[(k-1)/2] < nums[k] {
		nums[(k-1)/2], nums[k] = nums[k], nums[(k-1)/2]
		k = (k - 1) / 2
	}
}
