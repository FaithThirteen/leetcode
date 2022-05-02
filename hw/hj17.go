package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	strs := strings.Split(s, ";")
	x, y := 0, 0
	for i := 0; i < len(strs); i++ {
		pass := filter(strs[i])
		if !pass {
			continue
		}
		move := strs[i][0:1]                 // 移动键位
		step, _ := strconv.Atoi(strs[i][1:]) // 移动步数
		if move == "A" {
			x -= step
		}
		if move == "D" {
			x += step
		}
		if move == "W" {
			y += step
		}
		if move == "S" {
			y -= step
		}
		fmt.Printf("move: %s, step: %d\n", move, step)
	}
	//fmt.Printf("%d,%d\n",x,y)
}

func filter(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] != 'A' && s[i] != 'D' && s[i] != 'W' && s[i] != 'S' {
				return false
			}
		}
		if s[i] < '0' && s[i] > '9' {
			return false
		}
	}
	return true
}
