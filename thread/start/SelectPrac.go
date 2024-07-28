package SelectPrac

import (
	"fmt"
	"time"
)

func InitSelectPrac() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}
}

func Timeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func Default() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	for {
		select {
		case res := <-ch:
			fmt.Println("Received:", res)
			return
		case <-time.After(time.Millisecond * 500):
			fmt.Println("Timeout")
		default:
			fmt.Println("No channels ready, doing other work")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
