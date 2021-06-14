package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func die(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		die("Need a file to read")
	}

	fh, err := os.Open(os.Args[1])

	if err != nil {
		die(fmt.Sprintf("Cannot read file %s:%v", os.Args[1], err))
	}

	defer fh.Close()

	reader := csv.NewReader(fh)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'
	//reader.Comment = '#'
	i := 0

	for {
		i++
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Got error in line %d: %s", i, err)
		} else {
			for _, data := range record {
				fmt.Printf("- % s- ", data)
			}
		}

		fmt.Println("")
	}
}
