package channelAndContext

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func makeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func installTire(tireCh, painCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		painCh <- car
	}
	wg.Done()
	close(painCh)
}

func paintCar(painCh chan *Car) {
	for car := range painCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

func Factory() {
	tireCh := make(chan *Car)
	painCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go makeBody(tireCh)
	go installTire(tireCh, painCh)
	go paintCar(painCh)

	wg.Wait()
	fmt.Println("Close the factory")
}
