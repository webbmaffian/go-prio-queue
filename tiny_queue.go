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
	values     [tinyQueueMaxSize]*V
	prios      [tinyQueueMaxSize]P
	lastPopped *V
	start      uint8
	size       uint8
	compare    func(a, b P) bool
}

func (q *tinyQueue[V, P]) Size() uint64 {
	return uint64(q.size)
}

func (q *tinyQueue[V, P]) MaxSize() uint64 {
	return tinyQueueMaxSize
}

func (q *tinyQueue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *tinyQueue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *tinyQueue[V, P]) Pop() (value *V, prio P) {
	if q.size == 0 || q.values[q.start] == nil {
		return
	}

	q.values[q.start], q.lastPopped = q.lastPopped, q.values[q.start]

	value = q.lastPopped
	prio = q.prios[q.start]
	q.start = (q.start + 1) % tinyQueueMaxSize
	q.size--
	return
}

func (q *tinyQueue[V, P]) Push(value V, prio P) {
	if q.size == tinyQueueMaxSize && !q.compare(prio, q.prios[(q.start+q.size-1)%tinyQueueMaxSize]) {
		return
	}

	if q.size != tinyQueueMaxSize {
		q.size++
	}

	// Put value first in queue
	q.start = (q.start - 1) % tinyQueueMaxSize

	if q.values[q.start] == nil {
		q.values[q.start] = new(V)
	}

	*q.values[q.start] = value
	q.prios[q.start] = prio

	i := q.start

	for {
		j := (i + 1) % tinyQueueMaxSize

		if j == (q.start+q.size)%tinyQueueMaxSize || q.compare(q.prios[i], q.prios[j]) {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i = (i + 1) % tinyQueueMaxSize
	}
}
