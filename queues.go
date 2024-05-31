package queues

import (
	"fmt"
	"strings"
)

const initCapacity = 16

type Queue[T any] struct {
	ring  []T // length must be a power of two
	mask  int // must be a power of two minus 1
	start int
	end   int
	size  int
}

// New returns a new Queue.
func New[T any]() *Queue[T] {
	return NewWithCapa[T](initCapacity)
}

// NewWithCapa returns a new Queue with an internal capacity of at leat n.
func NewWithCapa[T any](n int) *Queue[T] {
	n = capa(n)
	return &Queue[T]{
		ring:  make([]T, n),
		mask:  int(n - 1),
		start: 0,
		end:   0,
		size:  0,
	}
}

func capa(n int) int {
	if n <= initCapacity {
		return initCapacity
	}
	if (n-1)&n == 0 { // is power of two?
		return n
	}
	// Compute the next power of two.
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Clear() {
	q.start = 0
	q.end = 0
	q.size = 0
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.size == len(q.ring) {
		q.resize()
	}
	q.ring[q.end] = elem
	q.end = (q.end + 1) & q.mask
	q.size++
}

func (q *Queue[T]) resize() {
	newRing := make([]T, len(q.ring)*2)
	if q.start == 0 {
		copy(newRing, q.ring)
		q.ring = newRing
		q.mask = len(newRing) - 1
		q.end = q.size
	} else {
		l := len(q.ring) - q.start
		copy(newRing[:l], q.ring[q.start:])
		copy(newRing[l:], q.ring[:q.end])
		q.start = 0
		q.end = len(q.ring)
		q.ring = newRing
		q.mask = len(newRing) - 1
	}
}

func (q *Queue[T]) Dequeue() T {
	if q.size == 0 {
		panic("pop on an empty queue")
	}
	elem := q.ring[q.start]
	q.start = (q.start + 1) & q.mask
	q.size--
	return elem
}

func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "Queue[]"
	}
	sb := strings.Builder{}
	sb.WriteString("Queue[")
	sb.WriteString(fmt.Sprintf("%v", q.ring[q.start]))
	for i := 1; i < q.Size(); i++ {
		p := (q.start + i) & q.mask
		sb.WriteString(fmt.Sprintf(" %v", q.ring[p]))
	}
	sb.WriteByte(']')
	return sb.String()
}
