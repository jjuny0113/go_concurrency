package SelectPrac

import (
	"fmt"
	"time"
)

func helloRunnable() {
	fmt.Println("goroutine: run()")
}

func Init() {
	for i := 0; i < 3; i++ {
		go helloRunnable()
	}
	fmt.Println("main: main() end")
	time.Sleep(time.Second)
}

func Count() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("value: %d\n", i)
		time.Sleep(time.Second)
	}
}

func InnerFunc() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("value: %d\n", i)
			time.Sleep(time.Second)
		}

	}()
	time.Sleep(6 * time.Second)
}

func printWork(content string, sleepMs time.Duration) {
	for {
		fmt.Println(content)
		time.Sleep(sleepMs)
	}
}

func PrintWorkExe() {
	go printWork("A", time.Second)
	go printWork("B", 500*time.Millisecond)
	select {}
}
