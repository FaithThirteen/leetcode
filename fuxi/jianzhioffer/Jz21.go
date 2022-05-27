package main

import "fmt"

// o(n),o(n)
func reOrderArray(array []int) []int {
	// write code here
	newArr := make([]int, 0, len(array))

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			newArr = append(newArr, array[i])
		}
	}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			newArr = append(newArr, array[i])
		}
	}

	return newArr
}

// o(n^2),o(n)
func reOrderArray1(array []int) []int {
	// write code here
	index := 0
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			// 将index到 i-1的位置往后移
			tmp := array[i]
			for k:=i-1;k>=index;k-- {
				array[k+1] = array[k]
			}
			array[index] = tmp
			index++
		}
	}
	return array
}

func main() {
	fmt.Println(reOrderArray([]int{1, 3, 5, 6, 7}))
	fmt.Println(reOrderArray1([]int{1, 3, 5, 6, 7}))
}
