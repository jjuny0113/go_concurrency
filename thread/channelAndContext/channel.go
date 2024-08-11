package channelAndContext

import (
	"fmt"
	"sync"
	"time"
)

var messages chan string = make(chan string)

// 채널에 데이터 넣기
//messages <- "this is a message"

// 채널에 데이터 빼기
//var msg string = <-messages

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

func Channel() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func SizeError() {
	ch := make(chan int)
	ch <- 9
	fmt.Println("Never print")
}

func butterSquare(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func Buffer() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go butterSquare(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func selectSquare(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func Select() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)
	wg.Add(1)
	go selectSquare(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true
	wg.Wait()
}

func tickSquare(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func Tick() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go tickSquare(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
