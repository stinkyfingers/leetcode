package leetcode

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type trie struct {
	target int
	combos [][]int
}

type node struct {
	value int
	prev  *node
}

func combinationSum(candidates []int, target int) [][]int {
	var output [][]int
	for i, candidate := range candidates {
		t := &trie{
			target: target,
		}
		n := &node{value: candidate}
		t.comboTrie(n, candidates[i:])
		output = append(output, t.combos...)
	}
	return output
}

func (t *trie) comboTrie(n *node, candidates []int) {
	if n.total() == t.target {
		t.combos = append(t.combos, n.values())
		return
	}
	if n.total() > t.target { // too high
		return
	}
	for i, candidate := range candidates {
		if candidate+n.value > t.target { // too high
			continue
		}
		current := &node{prev: n, value: candidate}
		t.comboTrie(current, candidates[i:])
	}

	return
}

// total calcs value of branch
func (n *node) total() int {
	var output int
	for _, value := range n.values() {
		output += value
	}
	return output
}

// values returns array of values of branch
func (n *node) values() []int {
	var output []int
	current := n
	for current != nil {
		output = append(output, current.value)
		current = current.prev
	}
	return output
}

/* testing */

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{1},
			target:     1,
			expected:   [][]int{{1}},
		},
		{
			candidates: []int{1},
			target:     2,
			expected:   [][]int{{1, 1}},
		},
		{
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
	}
	for _, test := range tests {
		res := combinationSum(test.candidates, test.target)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}

func (n *node) String() string {
	return strconv.Itoa(n.value)
}

func (n *node) Print() {
	fmt.Println(n.value)
}

func (t *trie) String() string {
	return fmt.Sprintf("%v", t.combos)
}
