package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}

func main() {
	a := make([]int, 5, 10)
	fmt.Println("len(a):", len(a))
	fmt.Println("cap(a):", cap(a))
	fmt.Println(a)

	b := make([]int,5)
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))
	fmt.Println(b)


}
