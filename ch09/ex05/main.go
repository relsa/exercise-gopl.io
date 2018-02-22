package main

import (
	"fmt"
	"time"
)

func main() {
	ch0 := make(chan struct{})
	ch1 := make(chan struct{})
	end := make(chan struct{})
	rally := make(chan int)

	go pingpong(ch0, ch1, end, rally)
	go pingpong(ch1, ch0, end, rally)

	ch0 <- struct{}{}

	timer := time.Tick(1 * time.Second)
	<-timer
	close(end)

	<-rally
	r := <-rally

	fmt.Printf("%d messages/sec\n", r)
}

func pingpong(in <-chan struct{}, out chan<- struct{}, end <-chan struct{}, rally chan<- int) {
	for i := 0; ; i++ {
		select {
		case v := <-in:
			select {
			case <-end:
				rally <- i
				return
			default:
				out <- v
			}
		case <-end:
			rally <- i
			return
		}
	}
}
