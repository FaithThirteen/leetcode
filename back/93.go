package main

import (
	"fmt"
	"strconv"
)

var result93 []string

// 与131题策略相同
func restoreIpAddresses(s string) []string {
	result93 = []string{}
	backtracking93(0, len(s), 0, s, "")
	return result93
}

/**
start:控制临时字符串的起始值，配合i列举出所有的可能
n:字符串大小
num: '.'出现的次数
s:源字符串
path:结果
*/
func backtracking93(start, n, num int, s, path string) {
	if num > 4 { // 大于四个，说明有五个段，过滤
		return
	}
	if start == len(s) {
		if num < 4 { // 采集结果，过滤掉不足四个段的ip
			return
		}
		result93 = append(result93, path[:len(path)-1])
	}

	for i := start; i < n; i++ {
		ip := s[start : i+1]
		if !isIp(ip) {
			continue
		}

		path = path + ip + "." // 每次都附加多一个，结尾进行处理
		backtracking93(i+1, n, num+1, s, path)
		path = path[:len(path)-1-len(ip)] // 回溯
	}

}

func isIp(s string) bool {
	if len(s) < 1 || len(s) > 3 { // 过滤空或者 '1111'
		return false
	}
	if s[0] == '0' && len(s) > 1 { // 过滤 '011'这类
		return false
	}
	n, _ := strconv.Atoi(s)
	if n > 255 { // 只能再0-255
		return false
	}

	return true
}

func main() {
	fmt.Println(restoreIpAddresses("101023"))

}

func restoreIpAddresses1(s string) []string {
	var res, path []string
	backTracking(s, path, 0, &res)
	return res
}
func backTracking(s string, path []string, startIndex int, res *[]string) {
	//终止条件
	if startIndex == len(s) && len(path) == 4 {
		tmpIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		*res = append(*res, tmpIpString)
	}
	for i := startIndex; i < len(s); i++ {
		//处理
		path := append(path, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			//递归
			backTracking(s, path, i+1, res)
		} else { //如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
		//回溯
		path = path[:len(path)-1]
	}
}
func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' { //对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
