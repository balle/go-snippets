package main

import "fmt"

func main() {
	var todos []string
	todos = make([]string, 5)

	dones := []string{"learn some go", "learn more go"}

	todos = append(todos, "learn even more go")
	fmt.Println(len(dones))
}
