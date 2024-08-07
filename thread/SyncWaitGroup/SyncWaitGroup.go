package SyncWaitGroup

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func DeferExam() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

	fmt.Println("Main Function")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func Main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, &wg)
	//go worker(2, &wg)
	wg.Wait()
	fmt.Println("All workers done")
}

func Wait() {
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		fmt.Printf("Goroutine start: NumGoroutine = %d\n", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutine end: NumGoroutine = %d\n", runtime.NumGoroutine())
	}()

	fmt.Printf("Main start: NumGoroutine = %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("Main end: NumGoroutine = %d\n", runtime.NumGoroutine())
}

func sumChan(start, end int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	fmt.Printf("Goroutine start for range %d to %d\n", start, end)
	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}

	result <- sum

}

func sum(start, end int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine start for range %d to %d\n", start, end)
	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}
	*result = sum
	fmt.Printf("Goroutine end for range %d to %d: sum = %d\n", start, end, sum)
}

func Sums() {
	var wg sync.WaitGroup
	result1, result2 := 0, 0
	wg.Add(2)
	go sum(1, 50, &result1, &wg)
	go sum(51, 100, &result2, &wg)
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	total := result1 + result2
	fmt.Printf("Total sum: %d\n", total)
}

func SumsWithChannel() {
	var wg sync.WaitGroup

	result := make(chan int, 2)
	wg.Add(2)
	go sumChan(1, 50, &wg, result)
	go sumChan(51, 100, &wg, result)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	close(result)
	sum1, sum2 := <-result, <-result
	total := sum1 + sum2
	fmt.Printf("Total sum: %d\n", total)
}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func Chanel() {
	ch := make(chan int)

	go producer(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func timeoutWorker(name string, duration time.Duration, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Printf("%s: Start\n", name)
	time.Sleep(duration)
	fmt.Printf("%s: Done\n", name)
	done <- true
}

func TimeoutMain() {
	var wg sync.WaitGroup
	done := make(chan bool, 1)

	wg.Add(1)
	go timeoutWorker("Worker1", 2*time.Second, &wg, done)
	select {
	case <-done:
		fmt.Println("Worker completed")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout waiting for worker")
	}
	wg.Wait()
	fmt.Println("Main: All done")
}

func sleep(duration time.Duration) {
	time.Sleep(duration)
}

func myTask(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	fmt.Printf("%s: Start\n", name)
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		sleep(1 * time.Second)
	}
	fmt.Printf("%s: Done\n", name)
}
func MainTask() {
	var wg sync.WaitGroup
	wg.Add(3)
	go myTask(&wg, "Worker1")
	go myTask(&wg, "Worker2")
	go myTask(&wg, "Worker3")
	wg.Wait()
	fmt.Println("모든 스레드 실행 완료")
}
