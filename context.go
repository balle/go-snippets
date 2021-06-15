package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func work(ctx context.Context, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()
	i := 1

	for {
		select {
		case <-ctx.Done():
			log.Println("Terminating go routine")
			return
		case out <- i:
			i++
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	defer close(ch)

	wg.Add(1)
	go work(ctx, &wg, ch)

	for x := 0; x < 3; x++ {
		fmt.Printf("Got %d\n", <-ch)
	}

	log.Println("Going to cancel")
	cancel()

	log.Println("Waiting")
	wg.Wait()

	log.Println("I am done.")
}
