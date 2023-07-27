# go-semaphore-vs-worker-pool

I used worker pool most of the time and found out there's also semaphore pattern as well.
Here's the example that I want to try to solve using both of these patterns.

## Description

We have some number of users from 1 to 100000. Every user should be deactivated using function Deactivate. It can be some API that does not support batch upload so we should send every user by its own ID separately. Every API request has its timeout which actually means the maximal time after which you'l have a response. Hence if you send all the requests to deactivate users in one sequence the maximal time time will be `users_count*timeout`. And the average time will circa this number as well.

We can decrease (make it circa `goroutines` times less) the total time by using the Golang concurrency and putting the request sending inside goroutines. Here come two strategies that I used to solve this:
- semaphore;
- worker pool.

Apart from showing the examples I did some benchmarks to compare these two patterns on the particular case mentioned above.

## Benchmark tests

First, you need to install `benchstat`.

Then run these two benchmark tests:
`go test -bench=DeactivateUsers ./workerpool -count 10  -run=^# -benchmem | tee workerpool.txt`
`go test -bench=DeactivateUsers ./semaphore -count 10  -run=^# -benchmem | tee semaphore.txt`

Then change the string `pkg: ...` in both text files so they're both identical and run this:
`benchstat semaphore.txt workerpool.txt`

And you'll get the results like the ones inside `benchstat.txt`