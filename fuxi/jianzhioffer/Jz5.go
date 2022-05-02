package main

import (
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {
	// write code here
	var ss strings.Builder
	for i:=0;i<len(s);i++{
		if s[i] != ' '{
			ss.WriteString(s[i:i+1])
		}else{
			ss.WriteString("%20")
		}
	}
	return ss.String()
}
