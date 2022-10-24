package prioqueue

import "math"

const smallQueueMaxSize = math.MaxUint8

func newSmallQueue[V any, P Number](compare func(a, b P) bool) Queue[V, P] {
	return &smallQueue[V, P]{
		compare: compare,
	}
}

type smallQueue[V any, P Number] struct {
	values     [256]*V
	prios      [256]P
	lastPopped *V
	start      uint8
	size       uint8
	compare    func(a, b P) bool
}

func (q *smallQueue[V, P]) Size() uint64 {
	return uint64(q.size)
}

func (q *smallQueue[V, P]) MaxSize() uint64 {
	return smallQueueMaxSize
}

func (q *smallQueue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *smallQueue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *smallQueue[V, P]) Pop() (value *V, prio P) {
	if q.size == 0 || q.values[q.start] == nil {
		return
	}

	q.values[q.start], q.lastPopped = q.lastPopped, q.values[q.start]

	value = q.lastPopped
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
