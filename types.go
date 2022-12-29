package prioqueue

// Any number.
type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Queue[V any, P Number] interface {

	// Get current size of the queue.
	Size() int

	// Get max size of the queue.
	MaxSize() int

	// Return whether queue is empty.
	Empty() bool

	// Reset queue for reuse.
	Reset()

	// Return a value without removing it from queue. Defaults to the first value. A negative index will begin
	// from the end (e.g. `Peek(-1)` will return the last value).
	Peek(idx ...int) (value V, prio P)

	// Remove and return the first value in the queue.
	Pop() (value V, prio P)

	// Remove and return the last value in the queue.
	PopLast() (value V, prio P)

	// Add value to queue.
	Push(value V, prio P)
}
