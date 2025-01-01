package trees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
	}
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (b *BinarySearchTree) Insert(value int) {
	newNode := NewNode(value)

	if b.Root == nil {
		b.Root = newNode
		return
	}

	node := b.Root
	for node != nil {
		if value < node.Value {
			if node.Left == nil {
				node.Left = newNode
				return
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = newNode
				return
			} else {
				node = node.Right
			}
		}
	}
}

func (b *BinarySearchTree) Lookup(value int) *Node {
	node := b.Root
	for node != nil {
		if value == node.Value {
			return node
		}
		if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
