package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func maxSlidingWindow0(nums []int, k int) []int {

	stack := make([]int, 0, len(nums))

	for i := range nums {
		l, r := i, i+k
		if r-1 < len(nums) {
			stack = max(nums[l:r], stack)
		}
	}
	return stack
}

func max(nums []int, stack []int) []int {

	for i, n := range nums {
		if i == 0 {
			stack = append(stack, n)
		} else {
			if stack[len(stack)-1] < n {
				stack[len(stack)-1] = n
			}
		}
	}
	return stack
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] } // 大根堆规则
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	res := make([]int, 1, n-k+1)
	res[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k { // 最大值不在区间内，从优先队列移除
			heap.Pop(q)
		}
		// 此时最大值肯定在区间内，将值存储
		res = append(res, nums[q.IntSlice[0]])
	}
	return res
}
