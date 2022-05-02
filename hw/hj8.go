package main

import (
	"fmt"
	"sort"
)

func main() {

	var count, num, value int
	fmt.Scan(&count)

	m := make(map[int]int, 0)
	arr := make([]int, 0)

	for i := 0; i < count; i++ {
		fmt.Scan(&num, &value)

		if _,ok := m[num];ok{ // 存储有，则arr不需要加入
			m[num] += value
			continue
		}
		m[num] += value
		arr = append(arr,num)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n",arr[i],m[arr[i]])
	}

}
