package turker

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumAToB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Println("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func WaitGroup() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sumAToB(i, 1000000000)
	}
	wg.Wait()
	fmt.Println("모든 계산이 완료됐습니다.")
}
