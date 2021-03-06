package structure

import "fmt"

type stackNode struct {
	next  *stackNode
	value interface{}
}

type LinkedListStack struct {
	top *stackNode // 栈顶节点
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (stack *LinkedListStack) IsEmpty() bool {
	if nil == stack.top {
		return true
	}
	return false
}

func (stack *LinkedListStack) Push(value interface{}) {
	stack.top = &stackNode{stack.top, value}
}

func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.top.value
	stack.top = stack.top.next
	return value
}

func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.top.value
}

func (stack *LinkedListStack) Flush() {
	stack.top = nil
}

func (stack *LinkedListStack) String() string {
	if stack.IsEmpty() {
		return "empty stack"
	} else {
		str := "top"
		for cur := stack.top; cur != nil; cur = cur.next {
			str += fmt.Sprintf(" --> %+v", cur.value)
		}
		return str + " --> tail"
	}
}
