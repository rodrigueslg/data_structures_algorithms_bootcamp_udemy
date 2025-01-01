package stacksqueues

type StackNode struct {
	Value any
	Prev  *StackNode
}

type Stack struct {
	Top    *StackNode
	Bottom *StackNode
	Length int
}

func NewStack() *Stack {
	return &Stack{
		Top:    nil,
		Bottom: nil,
		Length: 0,
	}
}

func NewStackNode(value any) *StackNode {
	return &StackNode{
		Value: value,
		Prev:  nil,
	}
}

func (s *Stack) Peek() any {
	if s.Top != nil {
		return s.Top.Value
	} else {
		return nil
	}
}

func (s *Stack) Push(value any) {
	newNode := NewStackNode(value)
	if s.Top == nil && s.Bottom == nil {
		s.Top = newNode
		s.Bottom = newNode
	} else {
		oldTop := s.Top
		s.Top = newNode
		s.Top.Prev = oldTop
	}
	s.Length++
}

func (s *Stack) Pop() any {
	var popped *StackNode
	if s.Top != nil {
		popped = s.Top
		s.Top = popped.Prev
		s.Length--
	} else {
		return nil
	}
	return popped.Value
}
