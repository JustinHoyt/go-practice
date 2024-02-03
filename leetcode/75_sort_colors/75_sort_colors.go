package main

// if the val is the
// Input:  [2,0,2,1,1,0]
func sortColorsMap(nums []int) {

	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	for i := range nums {
		switch {
		case counter[0] > 0:
			nums[i] = 0
			counter[0]--
		case counter[1] > 0:
			nums[i] = 1
			counter[1]--
		case counter[2] > 0:
			nums[i] = 2
			counter[2]--
		}
	}
}

/*
Input:  [0,2,2,1,1,0]
.........0............ zero
.........1............ one
.........2............ two
Output: [0,0,1,1,2,2]
*/
func sortColorsOnePass(nums []int) {
	idx0, idx1 := 0, 0

	for idx2, num := range nums {
		switch num {
		case 0:
			nums[idx2], nums[idx0] = nums[idx0], nums[idx2]
			if nums[idx1] > nums[idx2] {
				nums[idx2], nums[idx1] = nums[idx1], nums[idx2]
			}
			idx0++
			idx1++
		case 1:
			nums[idx2], nums[idx1] = nums[idx1], nums[idx2]
			idx1++
		}
	}
}

func main() {}
