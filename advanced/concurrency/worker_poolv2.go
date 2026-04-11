package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	processor()
}

type SendReportJob struct {
	ReportName string
	Format     string
	Size       int
}

func (srj *SendReportJob) processor() {
	fmt.Printf("Sending Report %s in %s of %d\n", srj.ReportName, srj.Format, srj.Size)
	time.Sleep(2 * time.Second)
	fmt.Println("Send Report Completed", srj.ReportName)
}

type SendEmailJob struct {
	Sub  string
	From string
	To   string
}

func (sej *SendEmailJob) processor() {
	fmt.Printf("Sending Email %s from %s to %s\n", sej.Sub, sej.From, sej.To)
	time.Sleep(2 * time.Second)
	fmt.Println("Send Report Completed", sej.Sub)
}

type WorkerPool struct {
	maxWorkers int
	Jobs       []Job
	jobsChan   chan Job
	wg         sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for job := range wp.jobsChan {
		job.processor()
	}
}

func (wp *WorkerPool) workerPoolExec() {
	wp.jobsChan = make(chan Job)

	wp.wg.Add(wp.maxWorkers)
	for i := 0; i < wp.maxWorkers; i++ {
		go wp.worker()
	}

	for _, job := range wp.Jobs {
		wp.jobsChan <- job
	}

	close(wp.jobsChan)
	wp.wg.Wait()
}

func WorkerPoolV2() {
	fmt.Println("-- Worker Pool V2 --")

	sej1 := &SendEmailJob{
		Sub:  "Your Order is created",
		From: "orders@ovw.com",
		To:   "customer@ovw.com",
	}

	sej2 := &SendEmailJob{
		Sub:  "Your Order is Processed",
		From: "orders@ovw.com",
		To:   "customer@ovw.com",
	}

	sej3 := &SendEmailJob{
		Sub:  "Your Order is Shipped",
		From: "orders@ovw.com",
		To:   "customer@ovw.com",
	}

	sej4 := &SendEmailJob{
		Sub:  "Your Order is Delivered",
		From: "orders@ovw.com",
		To:   "customer@ovw.com",
	}

	sej5 := &SendEmailJob{
		Sub:  "Your Order is Invoiced",
		From: "orders@ovw.com",
		To:   "customer@ovw.com",
	}

	srj1 := &SendReportJob{
		ReportName: "Sending report Downstream 1",
		Format:     "xlsx",
		Size:       25,
	}
	srj2 := &SendReportJob{
		ReportName: "Sending report Downstream 2",
		Format:     "xlsx",
		Size:       35,
	}
	srj3 := &SendReportJob{
		ReportName: "Sending report Downstream 3",
		Format:     "xlsx",
		Size:       45,
	}

	jobs := []Job{
		sej1,
		srj1,
		sej2,
		srj2,
		sej3,
		srj3,
		sej4,
		sej5,
	}

	workerPool := &WorkerPool{
		maxWorkers: 5,
		Jobs:       jobs,
	}
	workerPool.workerPoolExec()
}
