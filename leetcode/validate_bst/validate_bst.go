///usr/bin/true; exec /usr/bin/env go run "$0" "$@"

package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTRec(curr *TreeNode, leftBound int, rightBound int) bool {
	if curr == nil {
		return true
	}
	if leftBound >= curr.Val || curr.Val >= rightBound {
		return false
	}

	return isValidBSTRec(curr.Left, leftBound, curr.Val) &&
		isValidBSTRec(curr.Right, curr.Val, rightBound)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTRec(root, -math.MaxInt, math.MaxInt)
}

func main() {
	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   10,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 20},
		},
	}
	fmt.Println(isValidBST(root))
}
