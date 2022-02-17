package main

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	l list.List
}

func Constructor1() MyStack {
	return MyStack{
		l: list.List{},
	}
}

func (this *MyStack) Push(x byte) {
	this.l.PushFront(x)
}

func (this *MyStack) Pop() byte {
	v := this.l.Front().Value.(byte)
	this.l.Remove(this.l.Front())
	return v
}

func (this *MyStack) Top() byte {
	return this.l.Front().Value.(byte)
}

func (this *MyStack) Empty() bool {
	if this.l.Len() < 1 {
		return true
	}
	return false
}

func isValid(s string) bool {

	stack := Constructor1()
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	b := []byte(s)
	for _, b1 := range b {
		if b2,ok := m[b1];!ok{
			stack.Push(b1)
		}else{
			if !stack.Empty() && b2 == stack.Top() {
				stack.Pop()
				continue
			}
			return false
		}
	}

	return stack.Empty()
}

func main() {
	fmt.Println(isValid("]]]]"))
}
