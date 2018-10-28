package main

import (
	"flag"
)

const BUFFERSIZE = 1048

var filename, location string

func main() {
	flag.StringVar(&filename, "file", "", " Specific the file name to transfer")
	flag.StringVar(&location, "save", "", " Specific the location to save the file")
	flag.Parse()

	if filename != "" {
		Server()
	} else {
		if location == "" {
			println("File will be save at root OR use the -save Flag to specify a location")
		}
		Client()
	}
}
