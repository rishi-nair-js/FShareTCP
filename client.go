package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func Client() {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Println(" Connected to Server, start receiving the file name and fileSize")
	bufferFileSize := make([]byte, 10)
	bufferFileName := make([]byte, 64)
	connection.Read(bufferFileName)
	connection.Read(bufferFileSize)
	filelabel := strings.Trim(string(bufferFileName), ":")
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	newFile, err := os.Create(filelabel)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	for {
		var receivedBytes int64

		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, BUFFERSIZE-(fileSize-receivedBytes)))
			break
		}

		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	fmt.Println("File received completely!!")
}
