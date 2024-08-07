package turker

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func SeperatedArea() {
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.do()
			wg.Done()
		}()
	}
	wg.Wait()
}
