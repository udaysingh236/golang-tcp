package main

import (
	"fmt"
	"io"
	"net"
)

var (
	RequestQueue  = make(chan ClientRequest, 100)
	ResponseQueue = make(chan ClientResponse, 100)
)

func readNextByte(conn net.Conn, number uint32) ([]byte, bool) {
	bytes := make([]byte, number)

	_, err := io.ReadFull(conn, bytes)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return nil, false
	}
	if err != nil {
		panic(fmt.Sprintf("Error While Read %v", err))
	}
	return bytes, true

}

// CollectandValidate will collect and validate
func CollectandValidate(conn net.Conn) bool {
	buff, error := readNextByte(conn, 10)
	if !error {
		// fmt.Println("Something bad with Reading data maybe EOF")
		return true
	}
	RequestQueue <- ClientRequest{buff, conn}
	return true
}
