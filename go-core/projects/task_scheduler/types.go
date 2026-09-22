package main

import (
	"sync"
)

type Task struct {
	ID       string
	Name     string
	TaskType string
	wg       sync.WaitGroup
	stopChan chan struct{}
}

type Scheduler struct {
	tasks        map[string]*Task
	mu           sync.Mutex
	globalWG     sync.WaitGroup
	shuttingDown bool
}
