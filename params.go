package main

import (
	"flag"
	"fmt"
)

func main() {
	var processes = flag.Int("processes", 1, "number of processes")
	var url = flag.String("url", "", "url to fetch")
	var verbose = flag.Bool("verbose", false, "be verbose")
	flag.Parse()

	if *verbose == true {
		fmt.Printf("Fetching %s with %d processes\n", *url, *processes)
	}
}
