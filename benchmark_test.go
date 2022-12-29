package prioqueue

import (
	"crypto/rand"
	"strconv"
	"testing"
)

type dummyDynamicItem struct{}

func (d dummyDynamicItem) HigherThan(item DynamicQueueItem) bool {
	return true
}

func BenchmarkMinQueue_New(b *testing.B) {
	b.Run("Tiny_"+strconv.Itoa(tinyQueueMaxSize), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewTinyMinQueue[struct{}, uint8]()
		}
	})

	b.Run("Small_"+strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewSmallMinQueue[struct{}, uint8]()
		}
	})

	b.Run("Normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewMinQueue[struct{}, uint8](256)
		}
	})

	b.Run("Dynamic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = make(DynamicQueue, 0, 256)
		}
	})
}

func BenchmarkMinQueue_Push(b *testing.B) {
	b.Run("Tiny_"+strconv.Itoa(tinyQueueMaxSize), func(b *testing.B) {
		q := NewTinyMinQueue[struct{}, uint8]()
		prios := make([]byte, b.N)
		rand.Read(prios)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}
	})

	b.Run("Small_"+strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		q := NewSmallMinQueue[struct{}, uint8]()
		prios := make([]byte, b.N)
		rand.Read(prios)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}
	})

	b.Run("Normal", func(b *testing.B) {
		q := NewMinQueue[struct{}, uint8](256)
		prios := make([]byte, b.N)
		rand.Read(prios)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}
	})

	b.Run("Dynamic", func(b *testing.B) {
		q := make(DynamicQueue, 0, 256)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Push(dummyDynamicItem{})
		}
	})
}

func BenchmarkMinQueue_Pop(b *testing.B) {
	b.Run("Tiny_"+strconv.Itoa(tinyQueueMaxSize), func(b *testing.B) {
		q := NewTinyMinQueue[struct{}, uint8]()
		prios := make([]byte, b.N)
		rand.Read(prios)

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Pop()
		}
	})

	b.Run("Small_"+strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		q := NewSmallMinQueue[struct{}, uint8]()
		prios := make([]byte, b.N)
		rand.Read(prios)

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Pop()
		}
	})

	b.Run("Normal", func(b *testing.B) {
		q := NewMinQueue[struct{}, uint8](256)
		prios := make([]byte, b.N)
		rand.Read(prios)

		for i := 0; i < b.N; i++ {
			q.Push(struct{}{}, prios[i])
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Pop()
		}
	})

	b.Run("Dynamic", func(b *testing.B) {
		q := make(DynamicQueue, 0, 256)

		for i := 0; i < b.N; i++ {
			q.Push(dummyDynamicItem{})
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = q.Pop()
		}
	})
}

func BenchmarkMinQueue_Peek(b *testing.B) {
	b.Run("Tiny_"+strconv.Itoa(tinyQueueMaxSize), func(b *testing.B) {
		q := NewTinyMinQueue[struct{}, uint8]()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Peek()
		}
	})

	b.Run("Small_"+strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		q := NewSmallMinQueue[struct{}, uint8]()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Peek()
		}
	})

	b.Run("Normal", func(b *testing.B) {
		q := NewMinQueue[struct{}, uint8](256)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Peek()
		}
	})
}
