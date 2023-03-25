package utilities

// Queue : a container of objects (a linear collection) that are inserted and removed according to the first-in first-out (FIFO) principle
type Queue[T any] struct {
	data []T
}

// IsEmpty : checks whether the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Put : adds an element onto the queue and returns an pointer to the current queue
func (q *Queue[T]) Put(n T) {
	q.data = append(q.data, n)
}

// Get : removes the next element from the queue and returns its value
func (q *Queue[T]) Get() T {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}
