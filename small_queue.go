package prioqueue

import "math"

const smallQueueMaxSize = math.MaxUint8

func NewSmallQueue[V any, P Number](compare func(a, b P) bool) Queue[V, P] {
	return &smallQueue[V, P]{
		compare: compare,
	}
}

func NewSmallMinQueue[V any, P Number]() Queue[V, P] {
	return &smallQueue[V, P]{
		compare: Asc[P],
	}
}

func NewSmallMaxQueue[V any, P Number]() Queue[V, P] {
	return &smallQueue[V, P]{
		compare: Desc[P],
	}
}

type smallQueue[V any, P Number] struct {
	values  [smallQueueMaxSize + 1]*V
	prios   [smallQueueMaxSize + 1]P
	start   uint8
	size    uint8
	compare func(a, b P) bool
}

func (q *smallQueue[V, P]) Size() uint64 {
	return uint64(q.size)
}

func (q *smallQueue[V, P]) MaxSize() uint64 {
	return smallQueueMaxSize + 1
}

func (q *smallQueue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *smallQueue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *smallQueue[V, P]) Peek(idx ...uint64) (value V, prio P) {
	i := q.start

	if idx != nil {
		i += uint8(idx[0])
	}

	if q.values[i] == nil {
		return
	}

	return *q.values[i], q.prios[i]
}

func (q *smallQueue[V, P]) Pop() (value V, prio P) {
	if q.size == 0 || q.values[q.start] == nil {
		return
	}

	value = *q.values[q.start]
	prio = q.prios[q.start]
	q.start++
	q.size--
	return
}

func (q *smallQueue[V, P]) Push(value V, prio P) {
	if q.size == smallQueueMaxSize && !q.compare(prio, q.prios[q.start+q.size-1]) {
		return
	}

	if q.size != smallQueueMaxSize {
		q.size++
	}

	// Put value first in queue
	q.start--

	if q.values[q.start] == nil {
		q.values[q.start] = new(V)
	}

	*q.values[q.start] = value
	q.prios[q.start] = prio

	i := q.start

	for {
		j := i + 1

		if j == q.start+q.size || q.compare(q.prios[i], q.prios[j]) {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i++
	}
}
