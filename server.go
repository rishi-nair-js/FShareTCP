package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func Server() {
	// Create a TCP server.
	server, err := net.Listen("tcp", "localhost:27001")
	if err != nil {
		fmt.Println("Error Listening :", err)
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println(" Server Started ! Waiting for connection")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Client connected")
		go sendFileToClient(connection)

	}
}

func sendFileToClient(connection net.Conn) {
	defer connection.Close()
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileLabel := fillString(fileInfo.Name(), 64)
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fmt.Println("Sending Filename and Filesize")
	connection.Write([]byte(fileLabel))
	connection.Write([]byte(fileSize))
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println(" File has been send closing connection !!")
	os.Exit(0)
	return
}

func fillString(name string, Totalfill int) string {
	for {
		if len(name) < Totalfill {
			name = name + ":"
			continue
		}
		break
	}
	return name

}
