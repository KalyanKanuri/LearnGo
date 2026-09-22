package main

import (
	"time"
)

func main() {
	scheduler := NewScheduler()

	// runOnceTaskID := scheduler.RunOnce(
	// 	"test-task",
	// 	2*time.Second,
	// 	CheckAPIHealth,
	// )

	_ = scheduler.ScheduleTask(
		"repetetive-task",
		1*time.Second,
		2*time.Second,
		CheckAPIHealth,
	)

	time.AfterFunc(10*time.Second, scheduler.ShutDown)

	// Example of stopping a specific task
	// scheduler.StopTask(runOnceTaskID)
	scheduler.globalWG.Wait()
}
