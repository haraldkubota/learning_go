package medianSorted

import (
	"sort"
	// "fmt"
	"errors"
)

// Naive method: Merge arrays, find median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	res := 0.0
	if len(merged)%2 == 1 {
		res = float64(merged[(len(merged)-1)/2])
	} else {
		indexRight := len(merged) / 2
		res = float64(float64((merged[indexRight-1] + merged[indexRight])) / 2.0)
	}
	return res
}

// Faster:
// Walk both arrays until you reach the middle or what would be the merged array

// Get items from 2 sorted arrays
func getNext(i1, i2 int, nums1, nums2 []int) (int, int, int, error) {
	if i1+i2 >= len(nums1)+len(nums2) {
		return i1, i2, 0, errors.New("No more array member to use")
	}
	if i1 == len(nums1) {
		return i1, i2 + 1, nums2[i2], nil
	}
	if i2 == len(nums2) {
		return i1 + 1, i2, nums1[i1], nil
	}
	if nums1[i1] <= nums2[i2] {
		return i1 + 1, i2, nums1[i1], nil
	}

	return i1, i2 + 1, nums2[i2], nil
}

func findMedianSortedArraysFast(nums1 []int, nums2 []int) float64 {

	totalLength := len(nums1) + len(nums2)
	skipLength := 0
	if totalLength%2 == 1 {
		skipLength = (totalLength - 1) / 2
	} else {
		skipLength = totalLength/2 - 1
	}
	i1 := 0
	i2 := 0
	var val1, val2 int
	var res float64
	for i := 0; i < skipLength; i += 1 {
		i1, i2, _, _ = getNext(i1, i2, nums1, nums2)
	}
	i1, i2, val1, _ = getNext(i1, i2, nums1, nums2)
	if totalLength%2 == 1 {
		res = float64(val1)
	} else {
		i1, i2, val2, _ = getNext(i1, i2, nums1, nums2)
		res = float64(val1+val2) / 2.0
	}
	return res
}
