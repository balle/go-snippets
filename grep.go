package main

import ("os"
	"fmt"
	"log"
	"bufio"
	"regexp")

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "<string_to_grep>")
		os.Exit(1)
	}

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		for _, search := range os.Args {
			matched, err := regexp.Match(search, []byte(input.Text()))

			if err != nil {
				log.Fatal(err)
			}

			if matched {
				fmt.Println(input.Text())
				break
			}
		}
	}
}
