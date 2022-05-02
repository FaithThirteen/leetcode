package main

import "fmt"

func commonChars(words []string) []string {
	// 填充词频
	var fre [][]int
	for i := 0; i < len(words); i++ {
		row := make([]int, 26)
		for j := 0; j < len(words[i]); j++ {
			row[words[i][j]-'a']++
		}
		fre = append(fre, row)
	}

	// 遍历每列，
	var result []string
	for line := 0; line < 26; line++ {
		count := fre[0][line]
		for row := 0; row < len(fre); row++ {
			// 寻找最小
			count = min(count, fre[row][line])
		}
		s := string(rune(line + 'a'))
		for i := 0; i < count; i++ {
			result = append(result, s)
		}
	}
	return result
}



func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

	fmt.Println(commonChars([]string{"bella","label","roller"}))
}