package main

func checkPossibility(nums []int) bool {
	stat := 0
	lastInd := len(nums) - 1
	for prevI, v := range nums[1:] {
		// If the value of the current position is less
		// than the previous element value
		if nums[prevI] > v {
			// On the second occurrence, return false
			if stat == 1 {
				return false
			}
			// otherwise, mark the first
			stat++

			// 1. If the index of the previous element is not 0
			// 2. If the previous element of the current element is larger than the current element
			// or
			// 1. If the next index of the current index exists
			// 2. If the previous element is greater than the next element
			// If both conditions exist, return false
			if (prevI != 0 && nums[prevI-1] > v) &&
				(prevI+2 <= lastInd && nums[prevI] > nums[prevI+2]) {
				return false
			}
		}
	}

	return true
}
