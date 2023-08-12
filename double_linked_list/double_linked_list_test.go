package double_linked_list_test

import (
	"github.com/romulets/go-data-structures/double_linked_list"
	"testing"
)

func TestEnqueue(t *testing.T) {
	l := double_linked_list.New()

	isSizeExactly(t, l, 0)

	l.Enqueue(100)
	isSizeExactly(t, l, 1)

	l.Enqueue(97)
	isSizeExactly(t, l, 2)

	l.Enqueue(999866)
	isSizeExactly(t, l, 3)

	l.Enqueue(7887364827)
	l.Enqueue(123123)
	l.Enqueue(34534534)
	l.Enqueue(8098909)
	isSizeExactly(t, l, 7)
}

func TestDequeue(t *testing.T) {
	l := double_linked_list.New()
	l.Enqueue(8765)
	l.Enqueue(456)
	l.Enqueue(9876)
	isSizeExactly(t, l, 3)

	if n := l.Dequeue(); n != 8765 {
		t.Errorf("Expected to have dequeued %d, but it was %d", 8765, n)
	}
	isSizeExactly(t, l, 2)

	if n := l.Dequeue(); n != 456 {
		t.Errorf("Expected to have dequeued %d, but it was %d", 456, n)
	}
	isSizeExactly(t, l, 1)

	if n := l.Dequeue(); n != 9876 {
		t.Errorf("Expected to have dequeued %d, but it was %d", 9876, n)
	}
	isSizeExactly(t, l, 0)
}

func TestPop(t *testing.T) {
	l := double_linked_list.New()
	l.Enqueue(8765)
	l.Enqueue(456)
	l.Enqueue(9876)
	isSizeExactly(t, l, 3)

	if n := l.Pop(); n != 9876 {
		t.Errorf("Expected to have popped %d, but it was %d", 9876, n)
	}
	isSizeExactly(t, l, 2)

	if n := l.Pop(); n != 456 {
		t.Errorf("Expected to have popped %d, but it was %d", 456, n)
	}
	isSizeExactly(t, l, 1)

	if n := l.Pop(); n != 8765 {
		t.Errorf("Expected to have dequeued %d, but it was %d", 8765, n)
	}
	isSizeExactly(t, l, 0)
}

func TestLastFirst(t *testing.T) {
	l := double_linked_list.New()
	l.Enqueue(8765)
	l.Enqueue(456)
	l.Enqueue(9876)

	isFirstAndLastExactly(t, l, 8765, 9876)
	l.Pop()
	isFirstAndLastExactly(t, l, 8765, 456)
	l.Dequeue()
	isFirstAndLastExactly(t, l, 456, 456)
}

func isSizeExactly(t *testing.T, l double_linked_list.DoubleLinkedList, expected int) {
	t.Helper()

	if l.Size() != expected {
		t.Errorf("Expected list to have %d element(s), but it was %d", expected, l.Size())
	}
}
func isFirstAndLastExactly(t *testing.T, l double_linked_list.DoubleLinkedList, first int, last int) {
	t.Helper()

	if l.First() != first {
		t.Errorf("Expected to last to be %d, but it was %d", first, l.First())
	}

	if l.Last() != last {
		t.Errorf("Expected to last to be %d, but it was %d", last, l.Last())
	}
}
