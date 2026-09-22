# time pkg

## time.Timer

Timer also works as similar to sleep() in go, the difference timer gives us ability to
trigger use channels and execute code after specified duration, typically time.After uses
the same timer mechanism

### Use Cases

we use sleep when we need a simple delay, we can use timer when we need timeout logic, like
stopping the timer, resetting the timer and triggering an event after specified time based
on channel notification from timer.C

The main difference is sleep will block the goroutine completely it cannot process events
received on other channel whereas timer can achieve this this is something like putting
timer for doing a certain task whereas sleep is pausing completely for specified time.
