package prioqueue

// Create a new queue with 256 allocated values of V.
func NewQueue[V any, P number]() Queue[V, P] {
	return &queue[V, P]{}
}

type queue[V any, P number] struct {
	values [256]V
	prios  [256]P
	start  uint8
	length uint8
}

func (q queue[V, P]) Length() uint8 {
	return q.length
}

func (q *queue[v, P]) Reset() {
	q.start = 0
	q.length = 0
}

func (q *queue[V, P]) Pop() (value V, prio P) {
	if q.length == 0 {
		return
	}

	value = q.values[q.start]
	prio = q.prios[q.start]
	q.start++
	q.length--
	return
}

func (q *queue[V, P]) Push(value V, prio P) {
	if q.length == max && prio >= q.prios[q.start+q.length-1] {
		return
	}

	if q.length != max {
		q.length++
	}

	// Put value first in queue
	q.start--
	q.values[q.start] = value
	q.prios[q.start] = prio

	i := q.start

	for {
		j := i + 1

		if j == q.start+q.length || q.prios[i] < q.prios[j] {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i++
	}
}

func (q *queue[V, P]) PushReturnedValue(prio P, cb func(V) V) {
	if q.length == max && prio >= q.prios[q.start+q.length-1] {
		return
	}

	if q.length != max {
		q.length++
	}

	// Put value first in queue
	q.start--
	q.values[q.start] = cb(q.values[q.start])
	q.prios[q.start] = prio

	i := q.start

	for {
		j := i + 1

		if j == q.start+q.length || q.prios[i] < q.prios[j] {
			break
		}

		q.values[i], q.values[j] = q.values[j], q.values[i]
		q.prios[i], q.prios[j] = q.prios[j], q.prios[i]

		i++
	}

}
