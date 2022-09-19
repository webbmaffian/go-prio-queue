package prioqueue

import (
	"math/rand"
	"testing"
)

func BenchmarkPushSequential(b *testing.B) {
	q := NewQueue[struct{}, int]()
	s := struct{}{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(s, i)
	}
}

func BenchmarkPushRandom(b *testing.B) {
	q := NewQueue[struct{}, byte]()
	nums := make([]byte, b.N)
	rand.Read(nums)
	s := struct{}{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(s, nums[i])
	}
}

func BenchmarkPushSequentialStruct(b *testing.B) {
	type foobar struct {
		foo string
	}

	q := NewQueue[foobar, int]()
	var s foobar

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = foobar{foo: "bar"}
		q.Push(s, i)
	}
}

// func BenchmarkPushSequentialReturnedPointers(b *testing.B) {
// 	type foobar struct {
// 		foo string
// 	}

// 	q := NewQueue[*foobar, int]()

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		q.PushReturnedValue(i, func(f *foobar) *foobar {
// 			if f == nil {
// 				f = &foobar{}
// 			}

// 			f.foo = "baz"

// 			return f
// 		})
// 	}
// }

func BenchmarkPop(b *testing.B) {
	q := NewQueue[struct{}, int]()
	s := struct{}{}

	for i := 0; i < b.N; i++ {
		q.Push(s, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = q.Pop()
	}
}
