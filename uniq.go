package main

import ("os"
	"fmt"
	"bufio")

func main() {
	input := bufio.NewScanner(os.Stdin)
	lines := make(map[string]int)

	for input.Scan() {
		if lines[input.Text()] == 0 {
			lines[input.Text()] = 1
		}
	}

	for line,_ := range lines {
		fmt.Println(line)
	}
}
