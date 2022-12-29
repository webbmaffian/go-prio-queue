package prioqueue

func NewTinyQueue[V any, P Number](compare func(a, b P) bool) Queue[V, P] {
	return &tinyQueue[V, P]{
		compare: compare,
	}
}

func NewTinyMinQueue[V any, P Number]() Queue[V, P] {
	return &tinyQueue[V, P]{
		compare: Asc[P],
	}
}

func NewTinyMaxQueue[V any, P Number]() Queue[V, P] {
	return &tinyQueue[V, P]{
		compare: Desc[P],
	}
}

const tinyQueueMaxSize = 64

type tinyQueue[V any, P Number] struct {
	values  [tinyQueueMaxSize]*V
	prios   [tinyQueueMaxSize]P
	start   uint8
	size    uint8
	compare func(a, b P) bool
}

func (q *tinyQueue[V, P]) Size() int {
	return int(q.size)
}

func (q *tinyQueue[V, P]) MaxSize() int {
	return tinyQueueMaxSize
}

func (q *tinyQueue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *tinyQueue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *tinyQueue[V, P]) Peek(idx ...int) (value V, prio P) {
	i := q.start

	if idx != nil {
		i += uint8(idx[0])
	}

	i = q.wrap(i)

	if q.values[i] == nil {
		return
	}

	return *q.values[i], q.prios[i]
}

func (q *tinyQueue[V, P]) Pop() (value V, prio P) {
	if q.size == 0 || q.values[q.start] == nil {
		return
	}

	value = *q.values[q.start]
	prio = q.prios[q.start]
	q.start = q.wrap(q.start + 1)
	q.size--
	return
}

func (q *tinyQueue[V, P]) PopLast() (value V, prio P) {
	i := q.wrap(q.start + q.size - 1)

	if q.size == 0 || q.values[i] == nil {
		return
	}

	value = *q.values[i]
	prio = q.prios[i]
	q.size--
	return
}

func (q *tinyQueue[V, P]) Push(value V, prio P) {
	if q.size == tinyQueueMaxSize && !q.compare(prio, q.prios[q.wrap(q.start+q.size-1)]) {
		return
	}

	if q.size != tinyQueueMaxSize {
		q.size++
	}

	// Put value first in queue
	q.start = q.wrap(q.start - 1)

	if q.values[q.start] == nil {
		q.values[q.start] = new(V)
	}

	*q.values[q.start] = value
	q.prios[q.start] = prio

	i := q.start
	end := q.wrap(q.start + q.size)

	for {
		j := q.wrap(i + 1)

		if j == end || q.compare(q.prios[i], q.prios[j]) {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i = q.wrap(i + 1)
	}
}

func (q *tinyQueue[V, P]) wrap(i uint8) uint8 {
	return (i + tinyQueueMaxSize) % tinyQueueMaxSize
}
