package structure

import "testing"

func TestLinkedListStackPush(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s)
}

func TestLinkedListStackPop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s)

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s)
}

func TestLinkedListStackTop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
}
