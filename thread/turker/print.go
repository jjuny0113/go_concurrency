package turker

import (
	"fmt"
	"time"
)

func printHangul() {
	hangul := []string{"가", "나", "다", "라", "마", "바", "사"}
	for _, v := range hangul {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%v ", v)

	}
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func Print() {
	go printHangul()
	go printNumbers()

	time.Sleep(3 * time.Second)
}
