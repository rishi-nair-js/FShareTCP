package main

import (
	"flag"
)

const BUFFERSIZE = 1048

var filename string

func main() {
	flag.StringVar(&filename, "file", "", " Specific the file name to transfer")
	flag.Parse()

	if filename != "" {
		Server()
	} else {
		Client()
	}
}
