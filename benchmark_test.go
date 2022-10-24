package prioqueue

import (
	"math/rand"
	"strconv"
	"testing"
)

var testSizes = []uint64{16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384}

func BenchmarkSmallQueueNew(b *testing.B) {
	b.Run(strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = newSmallQueue[struct{}](Asc[byte])
		}
	})
}

func BenchmarkSmallQueuePush(b *testing.B) {
	b.Run(strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		q := newSmallQueue[struct{}](Asc[byte])
		nums := make([]byte, b.N)
		rand.Read(nums)
		s := struct{}{}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Push(s, nums[i])
		}
	})
}

func BenchmarkSmallQueuePop(b *testing.B) {
	b.Run(strconv.Itoa(smallQueueMaxSize), func(b *testing.B) {
		q := newSmallQueue[struct{}](Asc[byte])
		nums := make([]byte, b.N)
		rand.Read(nums)
		s := struct{}{}

		for i := 0; i < b.N; i++ {
			q.Push(s, nums[i])
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = q.Pop()
		}
	})
}

func BenchmarkCustomQueueNew(b *testing.B) {
	for _, size := range testSizes {
		b.Run(strconv.Itoa(int(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = newCustomQueue[struct{}, byte](Asc[byte], 16)
			}
		})
	}
}

func BenchmarkCustomQueuePush(b *testing.B) {
	for _, size := range testSizes {
		b.Run(strconv.Itoa(int(size)), func(b *testing.B) {
			q := newCustomQueue[struct{}, byte](Asc[byte], size)
			nums := make([]byte, b.N)
			rand.Read(nums)
			s := struct{}{}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				q.Push(s, nums[i])
			}
		})
	}
}

func BenchmarkCustomQueuePop(b *testing.B) {
	for _, size := range testSizes {
		b.Run(strconv.Itoa(int(size)), func(b *testing.B) {
			q := newCustomQueue[struct{}](Asc[byte], size)
			nums := make([]byte, b.N)
			rand.Read(nums)
			s := struct{}{}

			for i := 0; i < b.N; i++ {
				q.Push(s, nums[i])
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, _ = q.Pop()
			}
		})
	}
}
