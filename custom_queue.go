package prioqueue

func newCustomQueue[V any, P Number](compare func(a, b P) bool, maxSize uint64) Queue[V, P] {
	return &customQueue[V, P]{
		values:  make([]*V, maxSize),
		prios:   make([]P, maxSize),
		maxSize: maxSize,
		compare: compare,
	}
}

type customQueue[V any, P Number] struct {
	values     []*V
	prios      []P
	lastPopped *V
	start      uint64
	size       uint64
	maxSize    uint64
	compare    func(a, b P) bool
}

func (q *customQueue[V, P]) Size() uint64 {
	return q.size
}

func (q *customQueue[V, P]) MaxSize() uint64 {
	return q.maxSize
}

func (q *customQueue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *customQueue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *customQueue[V, P]) Pop() (value *V, prio P) {
	if q.size == 0 || q.values[q.start] == nil {
		return
	}

	q.values[q.start], q.lastPopped = q.lastPopped, q.values[q.start]

	value = q.lastPopped
	prio = q.prios[q.start]
	q.start = (q.start + 1) % q.maxSize
	q.size--
	return
}

func (q *customQueue[V, P]) Push(value V, prio P) {
	if q.size == q.maxSize && prio >= q.prios[(q.start+q.size-1)%q.maxSize] {
		return
	}

	if q.size != q.maxSize {
		q.size++
	}

	// Put value first in queue
	q.start = (q.start - 1) % q.maxSize

	if q.values[q.start] == nil {
		q.values[q.start] = new(V)
	}

	*q.values[q.start] = value
	q.prios[q.start] = prio

	i := q.start

	for {
		j := (i + 1) % q.maxSize

		// log.Println(i, j, q.start, q.size, q.maxSize)

		if j == (q.start+q.size)%q.maxSize || q.prios[i] < q.prios[j] {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i = (i + 1) % q.maxSize
	}
}
