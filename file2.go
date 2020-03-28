package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fh, err := os.Open("/etc/passwd")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		fmt.Println(strings.Split(scanner.Text(), ":")[0])
	}

	fh.Close()
}
