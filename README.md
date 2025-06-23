# goroutinePlayground
Learn and play with goroutine


=== Benchmark context switching

* go test  -bench=. -cpu=1 \taskSwitching_test.go

Output:

* goroutinePlayground % go test  -bench=. -cpu=1 \taskSwitching_test.go
* goos: darwin
* goarch: arm64
* cpu: Apple M2 Max
* BenchmarkContextSwitching       13768707                87.07 ns/op
* PASS
* ok      command-line-arguments  2.404s

-----------------------------------------------
== List of concurrency primitives

* WaitGroup
* Mutex/ RWMutex
* Once
* NewCond
* Pool