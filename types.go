package prioqueue

// Any number.
type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Queue[V any, P Number] interface {

	// Get current size of the queue.
	Size() uint64

	// Get max size of the queue.
	MaxSize() uint64

	// Return whether queue is empty.
	Empty() bool

	// Reset queue for reuse.
	Reset()

	// Remove and return the first value in the queue.
	Pop() (value *V, prio P)

	// Add value to queue.
	Push(value V, prio P)
}
