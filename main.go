package main

import "go_concurrency/thread/turker"

func main() {
	//fmt.Println("main: main() start")
	//go HelloThread.HelloGoroutine()
	//fmt.Println("main: main() 호출 후")
	//time.Sleep(time.Second)
	//fmt.Println("main: main() end")
	//ManyThreadMainV1.Init()
	//go ManyThreadMainV1.Count()
	//time.Sleep(6 * time.Second)

	//ManyThreadMainV1.InnerFunc()
	//ManyThreadMainV1.PrintWorkExe()
	//SelectPrac.InitSelectPrac()
	//SelectPrac.Timeout()
	//SelectPrac.Default()

	//SyncWaitGroup.DeferExam()
	//SyncWaitGroup.Main()
	//SyncWaitGroup.Wait()
	//SyncWaitGroup.MainTask()
	//context.ThreadStopMainV4()
	//context.PrinterV3()
	//context.Yield()
	//turker.Print()
	//turker.WaitGroup()
	//turker.DeadLock()
	turker.SeperatedArea()
}
