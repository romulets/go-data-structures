package double_linked_list

type DoubleLinkedList struct {
	head *linkedNode
	tail *linkedNode
	size int
}

type linkedNode struct {
	val  int
	next *linkedNode
	prev *linkedNode
}

func (q *DoubleLinkedList) Enqueue(val int) {
	n := &linkedNode{val: val, next: nil, prev: nil}

	if q.size == 0 {
		q.head = n
	}

	if q.size == 1 {
		n.prev = q.head
		q.head.next = n
		q.tail = n
	}

	if q.size > 1 {
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}

	q.size++
}

func (q *DoubleLinkedList) Dequeue() int {
	first := q.head
	q.head = q.head.next
	q.size--
	return first.val
}

func (q *DoubleLinkedList) Pop() int {
	var last *linkedNode

	// 1
	if q.size == 1 {
		last = q.head
		q.head = nil
		q.tail = nil

		// 1, 2
	} else if q.size == 2 {
		last = q.tail
		q.tail = nil

		// 1, 2, 3
	} else {
		last = q.tail
		q.tail = last.prev
	}

	q.size--
	return last.val
}

func (q *DoubleLinkedList) Last() int {
	if q.size == 1 {
		return q.head.val
	}

	return q.tail.val
}

func (q *DoubleLinkedList) First() int {
	return q.head.val
}
