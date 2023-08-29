package tree

type BinaryTreeNode[T interface{}] struct {
	val   T
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

func NewBinaryNode[T interface{}](val T) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (t *BinaryTreeNode[T]) Value() T {
	return t.val
}

func (t *BinaryTreeNode[T]) PutLeft(leftVal T) {
	t.left = NewBinaryNode(leftVal)
}

func (t *BinaryTreeNode[T]) PutRight(rightVal T) {
	t.right = NewBinaryNode(rightVal)
}

func (t *BinaryTreeNode[T]) Left() *BinaryTreeNode[T] {
	return t.left
}

func (t *BinaryTreeNode[T]) Right() *BinaryTreeNode[T] {
	return t.right
}
