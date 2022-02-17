package main

import "fmt"

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 

示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3
 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func main() {

	list := Constructor()
	FmtList(&list)
	fmt.Println("MyLinkedList Get: ", list.Get(4))

	list.AddAtHead(7) // 7为头节点
	fmt.Println("MyLinkedList AddAtHead:")
	FmtList(&list)

	list.AddAtTail(8)
	fmt.Println("MyLinkedList AddAtTail:")
	FmtList(&list)

	list.AddAtIndex(1,22)
	fmt.Println("MyLinkedList AddAtIndex:")
	FmtList(&list)

	list.DeleteAtIndex(6)
	fmt.Println("MyLinkedList DeleteAtIndex:")
	FmtList(&list)

}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

// 实例化链表
func InitList(lenth int) MyLinkedList {
	l := MyLinkedList{}
	head := &l
	for i := 0; i < lenth; i++ {
		head.Val = i
		head.Next = &MyLinkedList{}
		if i == lenth-1 {
			head.Next = nil
		}
		head = head.Next
	}
	return l
}
// 打印链表
func FmtList(list *MyLinkedList){
	head := list
	for head != nil {
		fmt.Println("val:", head.Val)
		head = head.Next
	}
	fmt.Println("-----------------------")
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {

	return InitList(4)
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	count := 0
	temp := this
	for temp != nil {
		if count == index {
			return temp.Val
		}
		temp = temp.Next
		count++
	}

	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	temp := *this
	this.Next = &temp
	this.Val = val
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	temp := this
	for temp != nil {
		if temp.Next == nil{
			temp.Next = &MyLinkedList{
				Val:  val,
				Next: nil,
			}
			return
		}
		temp = temp.Next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 小于0插在头部
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	// 等于放在尾部
	length := 0
	up := MyLinkedList{}
	up.Next = this
	cur := this
	for cur != nil{
		// 尾元素追加
		if cur.Next == nil &&length == index -1{
			cur.Next = &MyLinkedList{
				Val:  val,
				Next: nil,
			}
			return
		}else if cur.Next != nil && length == index{ // 小于则进行插入操作
			temp := &MyLinkedList{
				Val:  val,
				Next: nil,
			}
			up.Next = temp
			temp.Next = cur
			this.Next = up.Next
			return
		}
		up = *cur
		cur = cur.Next
		length++
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	length := 0
	up := &MyLinkedList{}
	up.Next = this
	cur := this.Next
	for cur != nil  {
		if length == index{
			up.Next = cur.Next
			this.Next = up.Next
			return
		}
		up = cur
		cur = cur.Next
		length++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
