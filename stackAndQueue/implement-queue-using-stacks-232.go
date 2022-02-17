package main

import (
	"container/list"
	"fmt"
)

/**
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


示例 1：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MyQueue struct {
	stack list.List
}

func Constructor() MyQueue {
	return MyQueue{
		stack: list.List{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack.PushBack(x)
}

func (this *MyQueue) Pop() int {
	v := this.stack.Front().Value.(int)
	this.stack.Remove(this.stack.Front())
	return v
}

func (this *MyQueue) Peek() int {
	return this.stack.Front().Value.(int)
}

func (this *MyQueue) Empty() bool {
	if this.stack.Len() < 1 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Peek() // 2
	param_2 := obj.Pop() // 2
	param_4 := obj.Empty() // false

	fmt.Printf("param_2: %d ,param_3: %d ,param_4: %v \n",param_2,param_3,param_4)
}
