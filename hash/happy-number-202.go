package main

import "fmt"

func main() {

	n := 2
	fmt.Println(isHappy(n))
}



func isHappy(n int) bool {

	m := make(map[int]struct{})

	for  {
		var sum = 0
		// 获取长度，对每一个数字进行平方累加
		l := length(n)
		for i := 0; i < l; i++ {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}

		n = sum
		if sum == 1 {
			return true
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}

	}

}

func length(num int) int {
	i := 0
	for num > 0 {
		num /= 10
		i++
	}
	return i
}
