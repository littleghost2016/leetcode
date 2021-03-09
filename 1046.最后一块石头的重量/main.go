package main

import (
	"fmt"
	"sort"
)

type myHeap struct {
	intSlice sort.IntSlice
}

func (h myHeap) Less(i, j int) bool {
	return h.intSlice[i] < h.intSlice[j]
}

func (h myHeap) Len() int {
	return len(h.intSlice)
}

// 没整明白这里为什么可以没有*，但下面Push和Pop两个函数必须有*
func (h *myHeap) Swap(i, j int) {
	h.intSlice[i], h.intSlice[j] = h.intSlice[j], h.intSlice[i]
}

func (h *myHeap) Push(i int) {
	h.intSlice = append(h.intSlice, i)
}

func (h *myHeap) Pop() int {
	temp := h.intSlice[len(h.intSlice)-1]
	h.intSlice = h.intSlice[:len(h.intSlice)-1]
	return temp
}

func lastStoneWeight(stones []int) int {
	myStones := &myHeap{stones}
	sort.Sort(myStones)

	for len(myStones.intSlice) > 1 {
		a := myStones.Pop()
		b := myStones.Pop()
		if (a - b) > 0 {
			myStones.Push(a - b)
		}
		sort.Sort(myStones)
		fmt.Println(myStones)
	}

	if len(myStones.intSlice) == 1 {
		return myStones.intSlice[0]
	} else {
		return 0
	}
}
