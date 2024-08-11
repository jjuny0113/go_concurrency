package channelAndContext

import (
	"context"
	"fmt"
	"time"
)

func printEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}

func Context() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 컨택스트 생성
	go printEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func SquareFunc() {
	wg.Add(1)
	//ctx := context.WithValue(context.Background(), "number", 9)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "keyword", "Lilly")
	go squareV2(ctx)
	cancel()
	wg.Wait()
}

func squareV2(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}
