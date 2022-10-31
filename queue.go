package prioqueue

func NewQueue[V any, P Number](compare func(a, b P) bool, maxSize uint64) Queue[V, P] {
	return &queue[V, P]{
		items:   make([]*queueItem[V, P], maxSize),
		maxSize: maxSize,
		compare: compare,
	}
}

func NewMinQueue[V any, P Number](maxSize uint64) Queue[V, P] {
	return &queue[V, P]{
		items:   make([]*queueItem[V, P], maxSize),
		maxSize: maxSize,
		compare: Asc[P],
	}
}

func NewMaxQueue[V any, P Number](maxSize uint64) Queue[V, P] {
	return &queue[V, P]{
		items:   make([]*queueItem[V, P], maxSize),
		maxSize: maxSize,
		compare: Desc[P],
	}
}

type queueItem[V any, P Number] struct {
	value V
	prio  P
}

type queue[V any, P Number] struct {
	items   []*queueItem[V, P]
	start   uint64
	size    uint64
	maxSize uint64
	compare func(a, b P) bool
}

func (q *queue[V, P]) Size() uint64 {
	return uint64(q.size)
}

func (q *queue[V, P]) MaxSize() uint64 {
	return q.maxSize
}

func (q *queue[V, P]) Empty() bool {
	return q.size == 0
}

func (q *queue[v, P]) Reset() {
	q.start = 0
	q.size = 0
}

func (q *queue[V, P]) Peek(idx ...uint64) (value V, prio P) {
	i := q.start

	if idx != nil {
		i += idx[0]
	}

	i = q.wrap(i)

	if q.items[i] == nil {
		return
	}

	return q.items[i].value, q.items[i].prio
}

func (q *queue[V, P]) Pop() (value V, prio P) {
	if q.size == 0 || q.items[q.start] == nil {
		return
	}

	value = q.items[q.start].value
	prio = q.items[q.start].prio
	q.start = q.wrap(q.start + 1)
	q.size--
	return
}

func (q *queue[V, P]) Push(value V, prio P) {
	if q.size == q.maxSize && !q.compare(prio, q.items[q.wrap(q.start+q.size-1)].prio) {
		return
	}

	if q.size != q.maxSize {
		q.size++
	}

	// Put value first in queue
	q.start = q.wrap(q.start - 1)

	if q.items[q.start] == nil {
		q.items[q.start] = new(queueItem[V, P])
	}

	q.items[q.start].value = value
	q.items[q.start].prio = prio

	i := q.start
	end := q.wrap(q.start + q.size)

	for {
		j := q.wrap(i + 1)

		if j == end || q.compare(q.items[i].prio, q.items[j].prio) {
			break
		}

		q.items[i], q.items[j] = q.items[j], q.items[i]

		i = q.wrap(i + 1)
	}
}

func (q *queue[V, P]) wrap(i uint64) uint64 {
	return (i + q.maxSize) % q.maxSize
}
