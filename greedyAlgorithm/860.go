package main

import "fmt"

func lemonadeChange(bills []int) bool {

	money := make(map[int]int, 3)
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			money[5]++
			continue
		} else if bills[i] == 10 {
			m, ok := money[5]
			if ok && m > 0 {
				money[5]--
				money[10]++
			} else {
				return false
			}

		} else if bills[i] == 20 {
			m1, _ := money[5]
			m2, _ := money[10]
			if m2 >= 1 && m1 >= 1 {
				money[5] = money[5] - 1
				money[10] = money[10] - 1
				money[20]++
				continue
			} else if m1 >= 3 {
				money[5] = money[5] - 3
				money[20]++
				continue
			} else {
				return false
			}
		}
	}

	return true
}

func lemonadeChange1(bills []int) bool {
	var five, ten int
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			five++
		case 10:
			five--
			ten++
		case 20:
			if ten > 0 {
				five--
				ten--
			} else {
				five -= 3
			}
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}
