package solver

type Queue[T any] struct {
	elem []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(e T) {
	q.elem = append(q.elem, e)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elem) > 0 {
		e := q.elem[0]
		q.elem = q.elem[1:]
		return e, true
	}
	var zeroVal T
	return zeroVal, false
}

func (q *Queue[T]) Len() int {
	return len(q.elem)
}
