package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//var arr []int
//
//type he struct {
//	sort.IntSlice
//}
//
//func(h he)Less(i,j int)bool{
//	return a[i] > a[j]
//}
//
//func (h *he)Push(v interface{}){
//	h.IntSlice = append(h.IntSlice,v.(int))
//}
//
//func (h *he)Pop()interface{}{
//	a:=h.IntSlice
//	v := a[len(a)-1]
//	h.IntSlice = a[:len(a)-1]
//	return v
//}

func topKFrequent1(nums []int, k int) []int {

	arr := []int{}
	m :=map[int]int{}

	// 放入map
	for _, n := range nums {
		m[n]++
	}
	for key,_ := range m {
		arr = append(arr,key)
	}

	sort.Slice(arr, func(i, j int) bool {
		return m[arr[i]] > m[arr[j]]
	})


	return arr[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3},2))
}

//方法一：小顶堆
func topKFrequent(nums []int, k int) []int {
	map_num:=map[int]int{}
	//记录每个元素出现的次数
	for _,item:=range nums{
		map_num[item]++
	}
	h:=&IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key,value:=range map_num{
		heap.Push(h,[2]int{key,value})
		if h.Len()>k{
			heap.Pop(h)
		}
	}
	res:=make([]int,k)
	//按顺序返回堆中的元素
	for i:=0;i<k;i++{
		res[k-i-1]=heap.Pop(h).([2]int)[0]
	}
	return res
}

//构建小顶堆
type IHeap [][2]int

func (h IHeap) Len()int {
	return len(h)
}

func (h IHeap) Less (i,j int) bool {
	return h[i][1]<h[j][1]
}

func (h IHeap) Swap(i,j int) {
	h[i],h[j]=h[j],h[i]
}

func (h *IHeap) Push(x interface{}){
	*h=append(*h,x.([2]int))
}
func (h *IHeap) Pop() interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[0:n-1]
	return x
}
