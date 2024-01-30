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
		res, ok := memo[remaining]
		switch {
		case ok:
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

func printTrie(node *Node, word string) {
	for char, childNode := range *node {
		if char == "S" {
			fmt.Printf("%s\n----------\n", word)
		} else if childNode != nil {
			printTrie(childNode, word+string(char))
		}
	}
}

type Node map[string]*Node

type Key struct {
	targetIdx int
	node      *Node
}

func buildTrie(words []string) *Node {
	start := &Node{}
	end := &Node{}

	for _, word := range words {
		curr := start
		for i := 0; i < len(word); i++ {
			char := string(word[i])
			if (*curr)[char] == nil {
				(*curr)[char] = &Node{}
			}
			curr = (*curr)[char]
		}
		(*curr)["S"] = start // Points back to the start
		(*curr)["E"] = end   // Points to the end node
	}

	return start
}

func wordBreakFiniteAutomata(target string, words []string) bool {
	start := buildTrie(words)

	memo := map[Key]bool{}
	var wordBreakRec func(int, *Node) bool
	wordBreakRec = func(currIdx int, curr *Node) bool {
		key := Key{currIdx, curr}
		if curr == nil {
			return false
		} else if cachedRes, cached := memo[key]; cached {
			return cachedRes
		} else if currIdx == len(target) && (*curr)["E"] != nil {
			return true
		} else if currIdx == len(target) {
			return false
		}

		char := string(target[currIdx])
		memo[key] = wordBreakRec(currIdx, (*curr)["S"]) || wordBreakRec(currIdx+1, (*curr)[char])
		return memo[key]
	}

	return wordBreakRec(0, start)
}

func main() {
	fmt.Println(wordBreakDP("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreakDP("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
