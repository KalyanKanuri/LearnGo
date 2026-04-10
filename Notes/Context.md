# Context

context pkg gives ability for go code to control a certain process like how much time it should live or until what time it should wait for a certain process, this is nothing but controlling timeout of the process which is waiting for another process.

lets say we have to do a db call for a long query for an api maybe user_details, for the api handler func we will set a timeout with context, in case if db is timedout or connection is not established we will not wait forever and cancel other tasks which are scheduled concurrently apart from db call so that we will not waste memory, cpu e.t.c., when we see the db connection failure with context, we will immediately cancel the remaining context and return from that process so that we can send response as timedout or anything we desired to the api.

## creating a context

```go
ctx := context.Context()
```

Context is an interface in pkg which contains Deadline(), Done(), Err(), Value() it's as below

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

* Deadline() gives the timeout which is configured for the current context
* Done() gives a channel indicating that the context is done it's work it can be done when it's errored out as well
* Err() gives the error happened for the context to be done
* Value() gives the value of provided key, this key and value is already passed through context when context is created as context.WithValue()

## Context Types

### Parent or Root Contexts

* context.Background() -> This creates a root or parent context with which we can create futher behavioural contexts, this is used mostly in prod systems.
* context.TODO() -> This is used a place holder when we are unsure of what context we should use, this is mostly unused.

### With series

* WithCancel(parent) -> provides a cancel() operation with which we can manually control the operations to stop.
* WithTimeout(parent, duration) -> Automatically cancels() after given duration
* WithDeadline(parent, timePoint) -> Automatically cancels() after given time point e.g., 5 PM.
* WithValue(parent, key, val) -> Used to pass metadata across functions
* WithoutCancel(parent) -> gives a copy of the parent which is not canceled even the parent is canceled, mostly used when the background tasks must complete even the original process fails.
