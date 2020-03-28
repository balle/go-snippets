package main

import ("os"
	"fmt"
	"log"
	"bufio")

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "<file>")
		os.Exit(1)
	}

	fh, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(fh)

	for input.Scan() {
		fmt.Println(input.Text())
	}

        fh.Close()
}
