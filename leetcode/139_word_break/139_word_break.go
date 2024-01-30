package main

import (
	"fmt"
	"regexp"
	"strings"
)

func wordBreakRegex(target string, words []string) bool {
	found, _ := regexp.MatchString(
		//          "^(cats|dog|sand|and|cat)*$"
		fmt.Sprintf("^(%s)*$", strings.Join(words, "|")),
		target,
	)
	return found
}

func wordBreakDP(target string, words []string) bool {
	memo := map[string]bool{}

	var wordBreakRec func(string) bool
	wordBreakRec = func(remaining string) bool {
		res, cached := memo[remaining]
		switch {
		case cached:
			return res
		case len(remaining) == 0:
			return true
		}

		for i := 0; i < len(words) && !res; i++ {
			if strings.HasPrefix(remaining, words[i]) {
				res = wordBreakRec(remaining[len(words[i]):])
			}
		}

		memo[remaining] = res
		return res
	}

	return wordBreakRec(target)
}

func wordBreakFiniteAutomata(target string, words []string) bool {
	// strings print better than runes for debugging
	type Node map[string]*Node
	type Key struct {
		targetIdx int
		node      *Node
	}

	buildTrie := func() *Node {
		start := &Node{}

		for _, word := range words {
			curr := start
			for i := 0; i < len(word); i++ {
				char := string(word[i])
				if (*curr)[char] == nil {
					(*curr)[char] = &Node{}
				}
				curr = (*curr)[char]
			}
			(*curr)["S"] = start
			(*curr)["E"] = &Node{}
		}

		return start
	}

	memo := map[Key]bool{}
	var wordBreakRec func(int, *Node) bool
	wordBreakRec = func(currIdx int, curr *Node) bool {
		key := Key{currIdx, curr}
		switch cacheRes, cached := memo[key]; {
		case cached:
			return cacheRes
		case curr == nil:
			return false
		case currIdx == len(target) && (*curr)["E"] != nil:
			return true
		case currIdx == len(target):
			return false
		}

		char := string(target[currIdx])
		memo[key] = wordBreakRec(currIdx, (*curr)["S"]) || wordBreakRec(currIdx+1, (*curr)[char])

		return memo[key]
	}

	start := buildTrie()
	return wordBreakRec(0, start)
}

func main() {}
