package context

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

//ThreadStopMainV4

func log(message string) {
	fmt.Println(message)
}

func sleep(duration time.Duration) {
	time.Sleep(duration)
}

func ThreadStopMainV4() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		myTask(ctx)
	}()

	sleep(100 * time.Millisecond)
	log("작업 중단 지시")
	cancel()
	wg.Wait() // go 루틴이 정리하고 끝날 때 까지 메인 고루틴 기다리는 중
	log("작업 완료")

}

func myTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log("작업정리")
			sleep(1 * time.Second)
			log("자원 종료")
			log("작업 종료")
			return
		default:
			log("작업중")
			time.Sleep(100 * time.Millisecond) // 작업 중
		}
	}
}

// Printer
type Printer struct {
	jobQueue chan string
	done     chan struct{}
}

func NewPrinter() *Printer {
	return &Printer{
		jobQueue: make(chan string, 100),
		done:     make(chan struct{}),
	}
}

func (p *Printer) Run() {
	for {
		select {
		case job := <-p.jobQueue:
			log("출력 시작: " + job)
			time.Sleep(3 * time.Second)
			log("출력 완료")
		case <-p.done:
			log("프린터 종료")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (p *Printer) AddJob(job string) {
	p.jobQueue <- job
}

func (p *Printer) Stop() {
	close(p.done)
}

func PrinterV3() {
	printer := NewPrinter()

	go printer.Run()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		log("프린터할 문서를 입력하세요. 종료(q) :")
		scanner.Scan()
		input := scanner.Text()
		if input == "q" {
			printer.Stop()
			log("프로그램 종료")
			return // 즉시 함수 종료
		}
		printer.AddJob(input)
	}
}

// yield
func Yield() {
	const threadCount = 1000
	var wg sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				fmt.Printf("Goroutine %d - %d\n", i, j)
				
				runtime.Gosched()
			}
		}(i)
	}

	wg.Wait()
}
