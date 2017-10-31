package main

import (
	"./io"
	"./servercli"
	"log"
	"strconv"
)

func main() {
	address, port, err := servercli.Init()

	if err {
		return
	}

	log.Println("Attempting to listen on", address+":"+strconv.Itoa(port))

	io.ListenForConnectionsForever(address, port)
}
