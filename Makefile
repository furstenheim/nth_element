bench:
	go test -v -bench=. -benchtime 5s
bench-graph:
	mkdir -p benchmarks/$$(git rev-parse HEAD)
	go test -run=XXX -bench . -cpuprofile benchmarks/$$(git rev-parse HEAD)/cpu.prof
	go tool pprof -svg FloydRivest.test benchmarks/$$(git rev-parse HEAD)/cpu.prof > benchmarks/$$(git rev-parse HEAD)/cpu.svg
