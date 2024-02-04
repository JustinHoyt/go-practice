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
	searchRec = func(slice []int, dir string) (res int, ok bool) {
		m := len(slice) / 2 // midpoint index
		switch {
		// Base cases
		case len(slice) == 1:
			return 0, target == slice[m]

		// Found
		case target == slice[m] && dir == "left":
			if res, ok := searchRec(slice[:m], dir); ok {
				return res, true
			}
			return m, true
		case target == slice[m] && dir == "right":
			if res, ok = searchRec(slice[m:], dir); ok {
				return res + m, true
			}
			return m, true

		// Search left
		case target < slice[m]:
			return searchRec(slice[:m], dir)

		// Search right
		default: // case target > slice[m]:
			res, ok := searchRec(slice[m:], dir)
			return res + m, ok // Add m to calculate offset from slice
		}
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
