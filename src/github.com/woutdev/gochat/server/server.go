package main

import (
	"./io"
	"./servercli"
	"fmt"
)

func main() {
	address, port, err := servercli.Init()

	if err {
		return
	}

	fmt.Println("Attempting to listen...")

	io.ListenForConnectionsForever(address, port)
}
