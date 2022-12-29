package prioqueue

import (
	"math"
	"sort"
)

type DynamicQueueItem interface {
	HigherThan(item DynamicQueueItem) bool
}

type DynamicQueue []DynamicQueueItem

func (q DynamicQueue) Size() int {
	return len(q)
}

func (q DynamicQueue) MaxSize() int {
	return math.MaxInt
}

func (q DynamicQueue) Empty() bool {
	return q == nil
}

func (q *DynamicQueue) Reset() {
	*q = (*q)[:0]
}

func (q DynamicQueue) Peek(idx ...int) DynamicQueueItem {
	if q == nil {
		return nil
	}

	if idx == nil || idx[0] == 0 {
		return q[q.findIndex(true)]
	} else if idx[0] == -1 {
		return q[q.findIndex(false)]
	}

	sort.Sort(q)

	return q[(idx[0]+len(q))%len(q)]
}

func (q *DynamicQueue) Pop() (item DynamicQueueItem) {
	if *q == nil {
		return nil
	}

	idx := q.findIndex(true)
	item = (*q)[idx]

	q.deleteIndex(idx)

	return
}

func (q *DynamicQueue) PopLast() (item DynamicQueueItem) {
	if *q == nil {
		return nil
	}

	idx := q.findIndex(false)
	item = (*q)[idx]

	q.deleteIndex(idx)

	return
}

func (q *DynamicQueue) Push(item DynamicQueueItem) {
	*q = append(*q, item)
}

func (q DynamicQueue) findIndex(highest bool) (idx int) {
	for i := range q {
		if i == 0 {
			continue
		}

		if q[i].HigherThan(q[idx]) == highest {
			idx = i
		}
	}

	return
}

func (q *DynamicQueue) deleteIndex(idx int) {
	if l := len(*q); l == 1 {
		q.Reset()
	} else {
		(*q)[idx], (*q)[l-1] = (*q)[l-1], (*q)[idx]
		*q = (*q)[:l-1]
	}
}

func (q DynamicQueue) Len() int {
	return len(q)
}

func (q DynamicQueue) Less(i, j int) bool {
	return q[i].HigherThan(q[j])
}

func (q DynamicQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
