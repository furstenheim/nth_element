bench:
	go test -v -bench=. -benchtime 5s
bench-graph:
	mkdir -p benchmarks/$$(date +%F)-$$(git rev-parse HEAD)
	go test -run=XXX -bench . -cpuprofile benchmarks/$$(date +%F)-$$(git rev-parse HEAD)/cpu.prof
	go tool pprof -svg FloydRivest.test benchmarks/$$(date +%F)-$$(git rev-parse HEAD)/cpu.prof > benchmarks/$$(date +%F)-$$(git rev-parse HEAD)/cpu.svg
