package lists

type DoublyLinkedListNode struct {
	Value any
	Prev  *DoublyLinkedListNode
	Next  *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head   *DoublyLinkedListNode
	Tail   *DoublyLinkedListNode
	Length int
}

func NewDoublyLinkedList(value any) *DoublyLinkedList {
	head := NewDoublyLinkedListNode(value, nil, nil)
	return &DoublyLinkedList{
		Head:   head,
		Tail:   head,
		Length: 1,
	}
}

func NewDoublyLinkedListNode(value any, prev *DoublyLinkedListNode, next *DoublyLinkedListNode) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Value: value,
		Prev:  prev,
		Next:  next,
	}
}

func (l *DoublyLinkedList) Append(value any) {
	newNode := NewDoublyLinkedListNode(value, l.Tail, nil)
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++
}

func (l *DoublyLinkedList) Prepend(value any) {
	oldHead := l.Head
	newHead := NewDoublyLinkedListNode(value, nil, oldHead)
	oldHead.Prev = newHead
	l.Head = newHead
	l.Length++
}

func (l *DoublyLinkedList) Insert(index int, value any) {
	if index < 0 || index > l.Length {
		panic("Index out of bounds")
	}
	leader := l.Get(index)
	follower := leader.Next

	newNode := NewDoublyLinkedListNode(value, leader, follower)
	leader.Next = newNode
	follower.Prev = newNode

	l.Length++
}

func (l *DoublyLinkedList) Remove(index int) {
	if index < 0 || index > l.Length {
		panic("Index out of bounds")
	}
	unwanted := l.Get(index)
	newNode := NewDoublyLinkedListNode(unwanted.Next.Value, unwanted.Prev, unwanted.Next.Next)
	unwanted.Value = newNode.Value
	unwanted.Prev = newNode.Prev
	unwanted.Next = newNode.Next
	l.Length--
}

func (l *DoublyLinkedList) Get(index int) *DoublyLinkedListNode {
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

func (l *DoublyLinkedList) ToArray() []any {
	values := []any{}
	node := l.Head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}
