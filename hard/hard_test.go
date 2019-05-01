package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type data struct {
		nums1 []int
		nums2 []int
	}

	testDatas := []data{{[]int{1, 3}, []int{2}}, {[]int{1, 2}, []int{3, 4}}, {[]int{1}, []int{1}}}
	except := []float64{2.0, 2.5, 1.0}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays(testData.nums1, testData.nums2) {
			t.Errorf("solution 1 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays2(testData.nums1, testData.nums2) {
			t.Errorf("solution 2 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays3(testData.nums1, testData.nums2) {
			t.Errorf("solution 3 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}
}