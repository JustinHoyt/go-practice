package main

/*
.0.1.2.3.4.5
[5,7,7,9]...
.L.....R....
...M........
............
*/
func searchRange(nums []int, target int) []int {
	var searchRec func([]int, string) (int, bool)
	searchRec = func(slice []int, dir string) (targetIdx int, ok bool) {
		midIdx := len(slice) / 2

		switch {
		// Base case
		case len(slice) == 1:
			targetIdx, ok = 0, target == slice[midIdx]

		// Found
		case target == slice[midIdx] && dir == "left":
			if targetIdx, ok = searchRec(slice[:midIdx], dir); !ok {
				targetIdx, ok = midIdx, true
			}
		case target == slice[midIdx] && dir == "right":
			targetIdx, ok = searchRec(slice[midIdx:], dir)
			targetIdx += midIdx
			if !ok {
				targetIdx, ok = midIdx, true
			}

		// Search left
		case target < slice[midIdx]:
			targetIdx, ok = searchRec(slice[:midIdx], dir)

		// Search right
		case target > slice[midIdx]:
			targetIdx, ok = searchRec(slice[midIdx:], dir)
			targetIdx += midIdx
		}

		return targetIdx, ok
	}

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftIdx, ok := searchRec(nums, "left")
	rightIdx, _ := searchRec(nums, "right")

	if ok {
		return []int{leftIdx, rightIdx}
	}
	return []int{-1, -1}
}
