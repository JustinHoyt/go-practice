package main

import (
	"slices"
)

/*
...........
[1,0,3,9,0]
...........
*/
func moveZeroes(nums []int) {
	var nonZeroIdx, zeroIdx int
	zeroIdx = slices.Index(nums, 0)
	if zeroIdx == -1 {
		return
	}
	for nonZeroIdx = zeroIdx; nonZeroIdx < len(nums); nonZeroIdx++ {
		if nums[nonZeroIdx] != 0 {
			break
		}
	}
	for zeroIdx < len(nums) && nonZeroIdx < len(nums) {
		nums[zeroIdx] = nums[nonZeroIdx]
		nums[nonZeroIdx] = 0

		zeroIdx++
		for nonZeroIdx < len(nums) && nums[nonZeroIdx] == 0 {
			nonZeroIdx++
		}
	}
}

func moveZeroesIter(nums []int) {
	for l := 0; l < len(nums); l++ {
		if nums[l] != 0 {
			continue
		}
		for r := l; r < len(nums); r++ {
			if nums[r] == 0 {
				continue
			}
			nums[l], nums[r] = nums[r], nums[l]
			break
		}
	}
}
