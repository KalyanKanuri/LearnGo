package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func genTaskID() string {
	return "task-" + uuid.NewString()
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

func (s *Scheduler) RunOnce(taskName string, delay time.Duration, action func()) string {
	t := Task{
		ID:       genTaskID(),
		Name:     taskName,
		TaskType: "Run-Once",
		stopChan: make(chan struct{}),
	}
	fmt.Printf("Scheduling %s - %s after %s time duration\n",
		t.ID,
		taskName,
		delay,
	)
	s.mu.Lock()
	if s.shuttingDown {
		s.mu.Unlock()
		return ""
	}
	s.tasks[t.ID] = &t
	s.globalWG.Add(1)
	s.mu.Unlock()
	t.wg.Add(1)

	select {
	case <-t.stopChan:
		fmt.Println("Task stopped before execution")
		t.wg.Done()
		s.globalWG.Done()
		return ""
	default:
		time.AfterFunc(delay, func() {
			defer t.wg.Done()
			defer s.globalWG.Done()
			select {
			case <-t.stopChan:
				return
			default:
			}
			fmt.Println("\n-- Executing task --")
			action()
			s.mu.Lock()
			delete(s.tasks, t.ID)
			s.mu.Unlock()
		})
	}
	return t.ID
}

func (s *Scheduler) ScheduleTask(
	taskName string, initDelay time.Duration,
	interval time.Duration, action func(),
) string {
	if interval <= 0 {
		fmt.Println("Invalid interval, must be greater than 0")
		return ""
	}

	t := Task{
		ID:       genTaskID(),
		Name:     taskName,
		TaskType: "recurring",
		stopChan: make(chan struct{}),
	}

	s.mu.Lock()
	if s.shuttingDown {
		s.mu.Unlock()
		return ""
	}
	s.tasks[t.ID] = &t
	s.globalWG.Add(1)
	s.mu.Unlock()

	t.wg.Add(1)

	go func() {
		defer t.wg.Done()
		defer s.globalWG.Done()

		// wait for given delay time
		initTimer := time.NewTimer(initDelay)
		defer initTimer.Stop()

		select {
		case <-initTimer.C:
			fmt.Printf("Starting task execution to run in %s intervals\n", interval)
		case <-t.stopChan:
			fmt.Println("scheduler stopped before executing initial run")
			return
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		fmt.Println("-- Executing recurring task --")
		for {
			select {
			case <-ticker.C:
				fmt.Println("Running task at ", time.Now().Format(time.RFC3339))
				action()
			case <-t.stopChan:
				return
			}
		}
	}()
	return t.ID
}

func (s *Scheduler) StopTask(taskID string) {
	s.mu.Lock()
	t, exists := s.tasks[taskID]
	if !exists {
		s.mu.Unlock()
		fmt.Printf("%s Task does not exists, exiting", taskID)
		return
	}

	delete(s.tasks, taskID)
	close(t.stopChan)
	s.mu.Unlock()
	t.wg.Wait()
}

func (s *Scheduler) ShutDown() {
	fmt.Println("Shutting down scheduler...")
	s.mu.Lock()
	if s.shuttingDown {
		s.mu.Unlock()
		return
	}
	s.shuttingDown = true
	tasks := make([]string, 0, len(s.tasks))
	for id := range s.tasks {
		tasks = append(tasks, id)
	}
	s.mu.Unlock()

	for _, id := range tasks {
		s.StopTask(id)
	}
	fmt.Println("All tasks stopped, scheduler shut down complete")
}
