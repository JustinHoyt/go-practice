package main

import (
	"testing"
)

type Test struct {
	name  string
	want  bool
	str   string
	words []string
}

type Solution struct {
	name string
	fn   func(string, []string) bool
}

func TestWordBreak(t *testing.T) {
	for _, sol := range []Solution{
		{"wordBreakRegex", wordBreakRegex},
		{"wordBreakDP", wordBreakDP},
		{"wordBreakFiniteAutomata", wordBreakFiniteAutomata},
		{"wordBreakGraph", wordBreakGraph},
	} {
		for _, tt := range []Test{
			{
				name:  "true case",
				want:  true,
				str:   "applepenapple",
				words: []string{"apple", "pen"},
			},
			{
				name:  "false case",
				want:  false,
				str:   "catsandog",
				words: []string{"cats", "dog", "sand", "and", "cat"},
			},
			{
				name:  "matches part of the last word",
				want:  false,
				str:   "aaaaaaa",
				words: []string{"aaaa", "aa"},
			},
			{
				name:  "matches twice in a row",
				want:  true,
				str:   "bb",
				words: []string{"a", "b", "bbb", "bbbb"},
			},
			{
				name:  "matches heavily branching trie",
				want:  false,
				str:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
				words: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			{
				name:  "matches twice in a row",
				want:  true,
				str:   "aaaaaaa",
				words: []string{"aaaa", "aaa"},
			},
			{
				name:  "matches on second prefix match",
				want:  true,
				str:   "goalspecial",
				words: []string{"go", "goal", "goals", "special"},
			},
		} {
			got := sol.fn(tt.str, tt.words)

			if got != tt.want {
				t.Errorf("%s(\n\t%s, \n\t%#v\n) \nWanted: %t\nGot: %t",
					sol.name,
					tt.str,
					tt.words,
					tt.want,
					got,
				)
			}
		}
	}
}
