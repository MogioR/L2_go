package circularqueue

type CircularQueue[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
	count  int
}

func NewCircularQueue[T any](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		buffer: make([]T, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (q *CircularQueue[T]) Push(item T) {
	if q.size == 0 {
		return
	}
	if q.count == q.size {
		q.head = (q.head + 1) % q.size
	} else {
		q.count++
	}
	q.buffer[q.tail] = item
	q.tail = (q.tail + 1) % q.size

}

func (q *CircularQueue[T]) Pull() T {
	if q.count == 0 {
		var zero T
		return zero
	}
	q.count--
	item := q.buffer[q.head]
	q.head = (q.head + 1) % q.size
	return item
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *CircularQueue[T]) Clear() {
	q.head = 0
	q.tail = 0
	q.count = 0
}

func (q *CircularQueue[T]) Len() int {
	return q.count
}
