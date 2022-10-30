// Fibonacci Heap
// Credits: https://github.com/theodesp/go-heaps
package prioqueue

func NewDynamicQueue[V any, P Number](compare func(a, b P) bool) Queue[V, P] {
	return &dynamicQueue[V, P]{
		compare: compare,
	}
}

func NewDynamicMinQueue[V any, P Number]() Queue[V, P] {
	return &dynamicQueue[V, P]{
		compare: Asc[P],
	}
}

func NewDynamicMaxQueue[V any, P Number]() Queue[V, P] {
	return &dynamicQueue[V, P]{
		compare: Desc[P],
	}
}

// FibonacciHeap is a implementation of Fibonacci heap.
type dynamicQueue[V any, P Number] struct {
	root    *node[V, P]
	size    uint64
	compare func(a, b P) bool
}

// node holds structure of nodes inside Fibonacci heap.
type node[V any, P Number] struct {
	value                     V
	prio                      P
	prev, next, parent, child *node[V, P]
	isMarked                  bool
	degree                    int
}

func (q *dynamicQueue[V, P]) Empty() bool {
	return q.root == nil
}

func (q *dynamicQueue[V, P]) Size() uint64 {
	return q.size
}

func (q *dynamicQueue[V, P]) MaxSize() uint64 {
	return 0
}

// Insert inserts a new node, with predeclared item, to the heap.
func (q *dynamicQueue[V, P]) Push(value V, prio P) {
	n := &node[V, P]{
		value:    value,
		prio:     prio,
		isMarked: false,
	}

	q.insertRoot(n)
}

// FindMin returns the minimum item.
func (q *dynamicQueue[V, P]) Peek() (value V, prio P) {
	if q.root == nil {
		return
	}

	return q.root.value, q.root.prio
}

// DeleteMin extracts the node with minimum item from a heap
// and returns the minimum item.
func (q *dynamicQueue[V, P]) Pop() (value V, prio P) {
	r := q.root
	if r == nil {
		return
	}
	for {
		// add r children to q's root list
		if x := r.child; x != nil {
			x.parent = nil
			if x.next != x {
				r.child = x.next
				x.next.prev = x.prev
				x.prev.next = x.next
			} else {
				r.child = nil
			}
			x.prev = r.prev
			x.next = r
			r.prev.next = x
			r.prev = x
		} else {
			break
		}
	}
	// remove r from q's root list
	r.prev.next = r.next
	r.next.prev = r.prev

	if r == r.next {
		q.root = nil
	} else {
		q.root = r.next
		q.consolidate()
	}

	q.size--

	return r.value, r.prio
}

func (q *dynamicQueue[V, P]) consolidate() {
	degreeToRoot := make(map[int]*node[V, P])
	w := q.root
	last := w.prev
	for {
		r := w.next
		x := w
		d := x.degree
		for {
			if y, ok := degreeToRoot[d]; !ok {
				break
			} else {
				if q.compare(y.prio, x.prio) {
					y, x = x, y
				}

				link(x, y)
				delete(degreeToRoot, d)
				d++
			}
		}
		degreeToRoot[d] = x
		if w == last {
			break
		}
		w = r
	}
	q.root = nil
	for _, v := range degreeToRoot {
		q.insertRoot(v)
	}

}

// Clear resets heap.
func (q *dynamicQueue[V, P]) Reset() {
	q.root = nil
}

func link[V any, P Number](x, y *node[V, P]) {
	// remove y from q's root list
	y.next.prev = y.prev
	y.prev.next = y.next
	// make y a child of x and increase degree of x
	y.parent = x
	if x.child == nil {
		x.child = y
		y.prev = y
		y.next = y
	} else {
		insert(x.child, y)
	}

	y.isMarked = false
}

func (q *dynamicQueue[V, P]) insertRoot(n *node[V, P]) {
	if q.root == nil {
		// create q's root list containing only n
		n.prev = n
		n.next = n
		q.root = n
	} else {
		// insert n to q's root list
		insert(q.root, n)

		if q.compare(n.prio, q.root.prio) {
			q.root = n
		}
	}

	q.size++
}

func insert[V any, P Number](x, y *node[V, P]) {
	x.prev.next = y
	y.next = x
	y.prev = x.prev
	x.prev = y
}
