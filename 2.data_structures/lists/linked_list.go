package lists

type LinkedListNode struct {
	Value any
	Next  *LinkedListNode
}

type LinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	Length int
}

func NewLinkedList(value any) *LinkedList {
	head := NewLinkedListNode(value, nil)
	return &LinkedList{
		Head:   head,
		Tail:   head,
		Length: 1,
	}
}

func NewLinkedListNode(value any, next *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{
		Value: value,
		Next:  next,
	}
}

func (l *LinkedList) Append(value any) {
	newNode := NewLinkedListNode(value, nil)
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++
}

func (l *LinkedList) Prepend(value any) {
	oldHead := l.Head
	newHead := NewLinkedListNode(value, oldHead)
	l.Head = newHead
	l.Length++
}

func (l *LinkedList) Insert(index int, value any) {
	if index < 0 || index > l.Length {
		panic("Index out of bounds")
	}
	node := l.Get(index)
	nodeCopy := NewLinkedListNode(node.Value, node.Next)
	node.Value = value
	node.Next = nodeCopy
	l.Length++
}

func (l *LinkedList) Remove(index int) {
	if index < 0 || index > l.Length {
		panic("Index out of bounds")
	}
	node := l.Get(index)
	node.Value = node.Next.Value
	node.Next = node.Next.Next
	l.Length--
}

func (l *LinkedList) Get(index int) *LinkedListNode {
	if index < 0 || index > l.Length {
		panic("Index out of bounds")
	}
	if index == 0 {
		return l.Head
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (l *LinkedList) ToArray() []any {
	values := []any{}
	node := l.Head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

func (l *LinkedList) Reverse() {
	if l.Head.Next == nil {
		return
	}
	first := l.Head
	second := first.Next
	l.Tail = l.Head
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	l.Head.Next = nil
	l.Head = first
}

//1, 10, 16, 88

// 1, 10, 1, 88
// 10, 10, 1, 88
// 10, 16, 1, 88

// 88, 16, 10, 1

func ReverseUsingPrepend(l LinkedList) *LinkedList {
	if l.Head == nil {
		panic("Null pointer operation")
	}
	reversed := NewLinkedList(l.Head.Value)
	node := l.Head.Next
	for node != nil {
		reversed.Prepend(node.Value)
		node = node.Next
	}
	return reversed
}
