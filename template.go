package main

import (
	"os"
	"text/template"
)

func main() {
	data := struct {
		Name string
	}{
		"Balle",
	}
	t, err := template.ParseFiles("template.txt")

	if err != nil {
		panic("Cannot parse template")
	}

	t.Execute(os.Stdout, data)
}
