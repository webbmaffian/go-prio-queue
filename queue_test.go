package prioqueue

import (
	"math/rand"
	"testing"
)

func BenchmarkPushSequential(b *testing.B) {
	q := NewQueue[struct{}, int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(struct{}{}, i)
	}
}

func BenchmarkPushRandom(b *testing.B) {
	q := NewQueue[struct{}, byte]()
	nums := make([]byte, b.N)
	rand.Read(nums)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(struct{}{}, nums[i])
	}
}

func BenchmarkPop(b *testing.B) {
	q := NewQueue[struct{}, int]()
	for i := 0; i < b.N; i++ {
		q.Push(struct{}{}, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
