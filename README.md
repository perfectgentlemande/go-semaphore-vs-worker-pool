# go-semaphore-vs-worker-pool
I want to bench and compare for my own educational purposes

Run this to get something:
`go test -bench=. -benchmem`

Or this if you're crazy:
`go test -bench=. ./workerpool -benchmem -benchtime=100000x`

Or this if you're mad:
`go test -bench=. ./semaphore -count 5 -benchmem | tee semaphore.txt`
`go test -bench=. ./workerpool -count 5 -benchmem | tee workerpool.txt`

Or this if you're insane:
`go test -bench=DeactivateUsers ./workerpool -count 10  -run=^# -benchmem | tee workerpool.txt`
`go test -bench=DeactivateUsers ./semaphore -count 10  -run=^# -benchmem | tee semaphore.txt`
