///usr/bin/true; exec /usr/bin/env go test .

package main

import (
	"encoding/json"
	"testing"
)

type Test struct {
	name string
	want bool
	root TreeNode
}


func TestValidateBST(t *testing.T) {
	tests := []Test{
		{
			name: "true case",
			root: TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   10,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 20},
				},
			},
			want: true,
		},
		{
			name: "false case",
			root: TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		got := isValidBST(&tt.root)

		str, _ := json.MarshalIndent(tt.root, "", "    ")

		if got != tt.want {
			t.Errorf("%s failed\nwant isValidBST(%s) = %+v\ngot = %+v", tt.name, str, tt.want, got)
		}
	}
}
