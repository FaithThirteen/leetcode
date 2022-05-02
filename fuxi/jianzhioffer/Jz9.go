package main

import "fmt"

var stack1 [] int
var stack2 [] int

func Push(node int) {
	stack1 = append(stack1,node)
}

func Pop() int{
	if len(stack2) == 0 { // 如果stack2临时队列没有元素了，开始填充
		for i:=len(stack1)-1;i>=0;i--{
			stack2 = append(stack2,stack1[i])
		}
		stack1 = []int{}
	}

	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}

func main() {
	Push(1)
	Push(2)
	fmt.Println(Pop())
	fmt.Println(Pop())
}