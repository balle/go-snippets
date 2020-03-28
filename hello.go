package main

import ( "bufio"
	 "strings"
	 "fmt"
	 "log"
	 "os" )

func main() {
        fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Hello,", strings.TrimSpace(name), "!")
	}
}
