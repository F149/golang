package main

import (
	"io"
	"os"
)

func main() {
	my_string := ""
	arguments := os.Args

	if len(arguments) == 1 {
		my_string = "Please gime one argument!"
	} else {
		my_string = arguments[1]
	}
	io.WriteString(os.Stdout, my_string)
	io.WriteString(os.Stdout, "\n")
}
