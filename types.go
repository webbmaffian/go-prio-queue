package prioqueue

import "math"

const max uint8 = math.MaxUint8

// Any number.
type number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Queue[V any, P number] interface {

	// Get the current length of the queue.
	Length() uint8

	// Reset the queue for reuse.
	Reset()

	// Remove and return the first value in the queue.
	Pop() (value V, prio P)

	// Add value to queue.
	Push(value V, prio P)

	// Add the returned value from callback to queue. Handy for juggling pointers.
	PushReturnedValue(prio P, cb func(V) V)
}
