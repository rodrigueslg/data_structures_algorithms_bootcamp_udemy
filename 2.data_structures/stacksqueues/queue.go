package stacksqueues

type QueueNode struct {
	Value any
	Next  *QueueNode
}

type Queue struct {
	First  *QueueNode
	Last   *QueueNode
	Length int
}

func NewQueue() *Queue {
	return &Queue{
		First:  nil,
		Last:   nil,
		Length: 0,
	}
}

func NewQueueNode(value any) *QueueNode {
	return &QueueNode{
		Value: value,
		Next:  nil,
	}
}

func (s *Queue) Peek() any {
	if s.First != nil {
		return s.First.Value
	} else {
		return nil
	}
}

func (s *Queue) Enqueue(value any) {
	newNode := NewQueueNode(value)
	if s.First == nil && s.Last == nil {
		s.First = newNode
		s.Last = newNode
	} else {
		oldLast := s.Last
		s.Last = newNode
		oldLast.Next = newNode
	}
	s.Length++
}

func (s *Queue) Dequeue() any {
	var dequeued *QueueNode
	if s.First != nil {
		dequeued = s.First
		s.First = dequeued.Next
		s.Length--
	} else {
		return nil
	}
	return dequeued.Value
}
