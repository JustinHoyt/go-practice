package main

/*
.................r.....
[1,3,9,0,0,0,0,0,1]
.......l...............
*/
func moveZeroesBrute(nums []int) {
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

/*
.n..........
[1,3,9,0,0]
.z..........
.0 1 2 3 4 5
*/
func moveZeroes(nums []int) {
	nextZero := func(idx int) int {
		for ; idx < len(nums) && nums[idx] != 0; idx++ {
		}
		return idx
	}
	nextNonZero := func(idx int) int {
		for ; idx < len(nums) && nums[idx] == 0; idx++ {
		}
		return idx
	}

	// Setup
	zeroIdx := nextZero(0)
	nonZeroIdx := nextNonZero(zeroIdx)

	// Iterative step
	for nonZeroIdx < len(nums) {
		nums[zeroIdx], nums[nonZeroIdx] = nums[nonZeroIdx], nums[zeroIdx]

		zeroIdx = nextZero(zeroIdx)
		nonZeroIdx = nextNonZero(zeroIdx)
	}
}
