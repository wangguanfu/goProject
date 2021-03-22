package main

import (
	"fmt"
	"math/rand"
)

// 任务放入chan
type Job struct {
	number int
	id     int
}
//结果放入chan
type Result struct {
	job    *Job
	sum int
}

func calc(job *Job, result chan *Result) {
	var sum int
	number := job.number
	for number != 0 {
		tmp := number % 10
		sum += tmp
		number /= 10
	}

	r := &Result{
		job:    job,
		sum: sum,
	}
	result <- r
}

func work(jobChan chan *Job, resultChan chan *Result) {
	for job:= range jobChan{
		calc(job, resultChan)
	}
}

func startWork(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go work(jobChan, resultChan)
	}
}

func printResult(resultChan chan *Result)  {
	for result:= range resultChan {
		fmt.Printf("job:%v, number:%v,result :%d\n",result.job.id, result.job.number,result.sum )
	}
}


func main() {

	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	startWork(128, jobChan, resultChan)

	go printResult(resultChan)

	var id int
	for {
		id ++
		number := rand.Int()
		job:=&Job{
			id :id,
			number: number,
		}
		jobChan <- job
	}
}
