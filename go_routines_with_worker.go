package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for _, fn := range files {
		wg.Add(1)

		go func() {
			// do something
			wg.Done()
		}()
	}

	wg.Wait()
}
