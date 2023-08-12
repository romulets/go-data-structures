package double_linked_list

type DoubleLinkedList[T interface{}] struct {
	head *linkedNode[T]
	tail *linkedNode[T]
	size int
}

type linkedNode[T interface{}] struct {
	val  T
	next *linkedNode[T]
	prev *linkedNode[T]
}

func New[T interface{}]() DoubleLinkedList[T] {
	return DoubleLinkedList[T]{size: 0, head: nil, tail: nil}
}

func (q *DoubleLinkedList[T]) Size() int {
	return q.size
}

func (q *DoubleLinkedList[T]) Enqueue(val T) {
	n := &linkedNode[T]{val: val, next: nil, prev: nil}

	switch q.size {
	case 0:
		q.head = n

	case 1:
		n.prev = q.head
		q.head.next = n
		q.tail = n

	default:
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}

	q.size++
}

func (q *DoubleLinkedList[T]) Dequeue() T {
	first := q.head
	q.head = q.head.next
	q.size--
	return first.val
}

func (q *DoubleLinkedList[T]) Pop() T {
	var last *linkedNode[T]

	switch q.size {
	case 1:
		last = q.head
		q.head = nil
		q.tail = nil

	case 2:
		last = q.tail
		q.tail = nil

	default:
		last = q.tail
		q.tail = last.prev
	}

	q.size--
	return last.val
}

func (q *DoubleLinkedList[T]) Last() T {
	if q.size == 1 {
		return q.head.val
	}

	return q.tail.val
}

func (q *DoubleLinkedList[T]) First() T {
	return q.head.val
}
