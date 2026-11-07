package queue

type node struct {
	val  int
	next *node
}

type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func (q *LinkedListQueue) Enqueue(val int) {
	newNode := &node{val: val}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *LinkedListQueue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	val := q.head.val
	q.head = q.head.next

	// that means last item is removed
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return val, true
}

func (q *LinkedListQueue) Peek() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	return q.head.val, true
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Size() int {
	return q.size
}
