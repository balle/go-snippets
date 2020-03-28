package main

func workloop(ch) {
	for fn := range ch {
		// do something
	}
}

func main() {
	ch := make(chan string, 2)

	for i := 0; i < 6; i++ {
		go workloop(ch)
	}

	for _, fn := range files {
		ch <- fn
	}
}
