package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// ClientRequest Will have input data
type ClientRequest struct {
	Message []byte
	Conn    net.Conn
}

// ClientResponse will have output data
type ClientResponse struct {
	Dataout []byte
	Conn    net.Conn
}

var (
	// Nworkers will have total number of workers
	Nworkers = flag.Int("worker", 4, "The number of workers you want to run..!")
)

//github.com/udaysingh236
func main() {
	StartDispatcher(*Nworkers)
	listener, err := net.Listen("tcp", "0.0.0.0:3333")
	if err != nil {
		fmt.Println("Error while Listening...!!", err)
		os.Exit(1)
	}
	fmt.Println("Server started listening....")
	defer listener.Close()
	///Receive Several conection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while Accepting...!!", err)
			os.Exit(1)
		}
		fmt.Println("Connection accepted...!!")
		go func() {
			for {
				if !CollectandValidate(conn) {
					break
				}
			}
			fmt.Println("Closing the conection..!!")
			conn.Close()
		}()

	}

}

func init() {

	flag.Parse()
}
