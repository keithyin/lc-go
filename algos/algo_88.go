package algos

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	fromPos := m - 1
	toPos := m + n - 1
	for fromPos >= 0 {
		nums1[toPos] = nums1[fromPos]
		fromPos--
		toPos--
	}
	nums1Cursor := toPos + 1
	nums2Cursor := 0

	insertPos := 0
	for nums1Cursor < (m+n) && nums2Cursor < n {
		if nums1[nums1Cursor] < nums2[nums2Cursor] {
			nums1[insertPos] = nums1[nums1Cursor]
			nums1Cursor++
		} else {
			nums1[insertPos] = nums2[nums2Cursor]
			nums2Cursor++
		}
		insertPos++
	}

	for nums2Cursor < n {
		nums1[insertPos] = nums2[nums2Cursor]
		nums2Cursor++
		insertPos++
	}

}
