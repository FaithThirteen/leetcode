package main

type MyStack struct {
	queue []int
}

func Constructor1() MyStack {
	return MyStack{queue: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	tempQ := this.queue
	this.queue = []int{x}
	this.queue = append(this.queue, tempQ...)
}

func (this *MyStack) Pop() int {
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

func (this *MyStack) Top() int {
	if len(this.queue) < 1 {
		return 0
	}
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	if len(this.queue) < 1 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type CQueue struct {
	stack1 []int // 原始栈
	stack2 []int // 临时栈
}

func Constructor2() CQueue {
	return CQueue{
		stack1: make([]int,0,10),
		stack2: make([]int,0,10),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) < 1 && len(this.stack1) > 0{
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = []int{}
	}

	if len(this.stack2) < 1{
		return -1
	}
	x := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return x
}


type MyQueue struct {
	stack1 []int
	stack2 []int
}


func Constructor3() MyQueue {
	return MyQueue{
		stack1: make([]int,0,10),
		stack2: make([]int,0,10),
	}
}


func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1,x)
}


func (this *MyQueue) Pop() int {
	if len(this.stack2) < 1 {
		for i:=len(this.stack1)-1;i>=0;i--{
			this.stack2 = append(this.stack2,this.stack1[i])
		}
		this.stack1 = []int{}
	}
	x := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return x
}


func (this *MyQueue) Peek() int {
	if len(this.stack2) < 1 {
		for i:=len(this.stack1)-1;i>=0;i--{
			this.stack2 = append(this.stack2,this.stack1[i])
		}
		this.stack1 = []int{}
	}
	return this.stack2[len(this.stack2)-1]
}


func (this *MyQueue) Empty() bool {
	if len(this.stack1) < 1 && len(this.stack2) < 1{
		return true
	}
	return false
}
