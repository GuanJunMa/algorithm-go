package structure

import "fmt"

type queueNode struct {
	val  interface{}
	next *queueNode
}

type LinkedListQueue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (queue *LinkedListQueue) EnQueue(value interface{}) {
	node := &queueNode{value, nil}
	if nil == queue.tail {
		queue.tail = node
		queue.head = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
	queue.length++
}

func (queue *LinkedListQueue) DeQueue() interface{} {
	if queue.head == nil {
		return nil
	}
	v := queue.head.val
	queue.head = queue.head.next
	queue.length--
	return v
}

func (queue *LinkedListQueue) String() string {
	if queue.head == nil {
		return "empty queue"
	} else {
		result := "head"
		for cur := queue.head; cur != nil; cur = cur.next {
			result += fmt.Sprintf(" --> %+v", cur.val)
		}
		result += "--> tail"
		return result
	}
}
