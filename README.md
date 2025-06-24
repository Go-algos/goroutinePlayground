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


-------------------------------------------------

Benchmark network demon:

go test -run BenchmarkNetworkRequest --benchtime=10s -bench=.

* go test -run BenchmarkNetworkRequest --benchtime=10s -bench=.

Accept connection from 127.0.0.1:50350
goos: darwin
goarch: arm64
pkg: goroutinePlayground
cpu: Apple M2 Max
BenchmarkNetworkRequest-12              Accept connection from 127.0.0.1:50351
Accept connection from 127.0.0.1:50352
Accept connection from 127.0.0.1:50353
Accept connection from 127.0.0.1:50354
Accept connection from 127.0.0.1:50355
Accept connection from 127.0.0.1:50356
Accept connection from 127.0.0.1:50357
Accept connection from 127.0.0.1:50358
Accept connection from 127.0.0.1:50359
Accept connection from 127.0.0.1:50360
10        1001897833 ns/op
BenchmarkContextSwitching-12            79680679               133.6 ns/op
PASS
ok      goroutinePlayground     21.987s
