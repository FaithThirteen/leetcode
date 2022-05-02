package main

import (
	"fmt"
)

func StrSplit(str string) []string {
	var newStrS []string
	var zeroByte = []byte("0")[0]

	for {
		// 尾部无字符直接退出分隔循环
		if len(str) == 0 {
			break
		}
		// 尾部不足八个字符需要补足
		if len(str) < 8 {
			// 要补足的字符
			var addByte = make([]byte, 8-len(str))
			for pos, _ := range addByte {
				addByte[pos] = zeroByte
			}
			newStrS = append(newStrS, fmt.Sprintf("%s%s", str, string(addByte)))
			break
		}
		// 切分8位的字符串
		newStrS = append(newStrS, str[:8])
		str = str[8:]
	}
	return newStrS
}

func main() {
	var str string
	fmt.Scan(&str)

	for{
		if len(str) == 0 {
			break
		}

		if len(str) < 8 {
			addByte := make([]byte,8-len(str))
			for i:=0;i<8-len(str);i++{
				addByte[i] =[]byte("0")[0]
			}
			fmt.Printf("%s%s\n",str,string(addByte))
			break
		}
		fmt.Printf("%s\n",str[:8])
		str = str[8:]
	}
}
