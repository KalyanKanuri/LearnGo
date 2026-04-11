# Worker Pool Pattern in Go

Worker Pool is a concurrency pattern in Go where we create a limited number of goroutines and reuse them for new tasks without creating new goroutine for each task, the tasks wait in Queue essentialy channels in Go

## Core Idea

* A channel to collect all the tasks [Tasks or Jobs]
* A func to process all the task as a stream
* A channel to collect the output [Results] (optional)
