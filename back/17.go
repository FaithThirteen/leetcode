package main

import "fmt"

var result []string

func letterCombinations(digits string) []string {
	result = []string{}
	if len(digits) < 1 {
		return result
	}
	var m = [10]string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	backtracking17(digits,"",0,m)
	return result
}

func backtracking17(digits,track string, index int, digitsMap [10]string) {
	if len(track) == len(digits) {
		result = append(result, track)
		return
	}
	tmpK := digits[index] - '0'  // 确定字符
	arr := digitsMap[tmpK]    // 取字符对应的字符串

	for i := 0;i <len(arr);i++{
		backtracking17(digits,track+string(arr[i]),index+1,digitsMap)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
