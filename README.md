# Go Prio Queue
A highly optimized [priority queue](https://en.wikipedia.org/wiki/Priority_queue) with generics. See [benchmark](#benchmark).
- 100% Go.
- Zero dependencies.
- Zero allocations for push/pop operations.
- Around 3 ns for push (up to 4000 values).
- Around 2 ns for pop.

## Install
```
go get github.com/webbmaffian/go-prio-queue
```

## Usage

### Min heap (ascending priority queue)
```go
q := NewMinQueue[string, float64]()

q.Push("a", 56)
q.Push("b", 12)
q.Push("c", 34)

for !q.Empty() {
    log.Println(q.Pop())
}

// b 12
// c 34
// a 56
```

### Max heap (descending priority queue)
```go
q := NewMaxQueue[string, float64]()

q.Push("a", 56)
q.Push("b", 12)
q.Push("c", 34)

for !q.Empty() {
    log.Println(q.Pop())
}

// a 56
// c 34
// b 12
```

## Benchmark
```
goos: darwin
goarch: arm64
pkg: github.com/webbmaffian/go-prio-queue
BenchmarkSmallQueueNew/255-10                 4459081           269.000 ns/op        2688 B/op           1 allocs/op
BenchmarkSmallQueuePush/255-10              349781355             3.427 ns/op           0 B/op           0 allocs/op
BenchmarkSmallQueuePop/255-10               586978732             2.041 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueueNew/16-10                14742232            75.690 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/32-10                15763598            75.530 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/64-10                15500386            75.620 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/128-10               16165234            75.990 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/256-10               15374652            75.800 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/512-10               16020327            76.480 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/1024-10              15808869            76.040 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/2048-10              15815224            76.110 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/4096-10              14650374            75.960 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/8192-10              16067761            76.890 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueueNew/16384-10             16040492            78.120 ns/op         240 B/op           3 allocs/op
BenchmarkCustomQueuePush/16-10              441543746             2.720 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/32-10              438644319             2.734 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/64-10              433402447             2.770 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/128-10             427948100             2.806 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/256-10             429108579             2.770 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/512-10             427958403             2.761 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/1024-10            389328944             2.797 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/2048-10            386863712             2.961 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/4096-10            298121041             3.472 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/8192-10             58608999            17.650 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePush/16384-10              114636         11915.000 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/16-10               573327946             2.063 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/32-10               583133943             2.043 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/64-10               581798670             2.059 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/128-10              583278972             2.049 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/256-10              581352746             2.062 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/512-10              582200670             2.056 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/1024-10             577774112             2.056 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/2048-10             578327653             2.057 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/4096-10             583535660             2.059 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/8192-10             584634346             2.062 ns/op           0 B/op           0 allocs/op
BenchmarkCustomQueuePop/16384-10            583455508             2.061 ns/op           0 B/op           0 allocs/op
```