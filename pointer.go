package main

import (
	"fmt"
)

func foo(numberp *int) {
	fmt.Printf("%x: %d\n", numberp, *numberp)
}

func main() {
	var i int = 42
	foo(&i)
}
