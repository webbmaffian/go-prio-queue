package prioqueue

// Create a new queue with `maxSize` allocated values of V. Defaults to a small queue of 256 values.
// The `compare` function should return true if a has higher priority than b.
func NewQueue[V any, P Number](compare func(a, b P) bool, maxSize ...uint64) Queue[V, P] {
	if len(maxSize) == 0 || maxSize[0] == 0 || maxSize[0] == smallQueueMaxSize {
		return newSmallQueue[V](compare)
	}

	return newCustomQueue[V](compare, maxSize[0])
}

// Min (ascending) priority queue.
func NewMinQueue[V any, P Number](maxSize ...uint64) Queue[V, P] {
	if len(maxSize) == 0 || maxSize[0] == 0 || maxSize[0] == smallQueueMaxSize {
		return newSmallAscQueue[V, P]()
	}

	return newCustomQueue[V](Asc[P], maxSize[0])
}

// Max (descending) priority queue.
func NewMaxQueue[V any, P Number](maxSize ...uint64) Queue[V, P] {
	if len(maxSize) == 0 || maxSize[0] == 0 || maxSize[0] == smallQueueMaxSize {
		return newSmallDescQueue[V, P]()
	}

	return newCustomQueue[V](Desc[P], maxSize[0])
}

func NewMinTinyQueue[V any, P Number]() Queue[V, P] {
	return newTinyAscQueue[V, P]()
}

// Max (descending) priority queue.
func NewMaxTinyQueue[V any, P Number](maxSize ...uint64) Queue[V, P] {
	return newTinyDescQueue[V, P]()
}
