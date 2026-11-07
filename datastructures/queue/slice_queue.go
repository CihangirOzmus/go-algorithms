package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return -1, false
	}

	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return -1, false
	}
	return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
