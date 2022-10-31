package prioqueue

import (
	"testing"
)

const size = 35_000

// func BenchmarkFindIterate(b *testing.B) {
// 	var haystack [size]byte

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		haystack[i%size] = 1

// 		for _, x := range haystack {
// 			if x == 1 {
// 				break
// 			}
// 		}

// 		haystack[i%size] = 0
// 	}
// }

// func BenchmarkFindBytealg(b *testing.B) {
// 	var haystack [size]byte

// 	b.ResetTimer()

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		haystack[i%size] = 1
// 		_ = bytes.IndexByte(haystack[:], 1)
// 		haystack[i%size] = 0
// 	}
// }

func BenchmarkBubbleMove(b *testing.B) {
	var data [size]byte
	data[0] = 1
	b.SetBytes(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for x := range data[:size-1] {
			data[x], data[x+1] = data[x+1], data[x]
		}
	}
}

func BenchmarkCopy(b *testing.B) {
	var data [size]byte
	data[0] = 1
	b.SetBytes(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(data[1:], data[:])
	}
}
