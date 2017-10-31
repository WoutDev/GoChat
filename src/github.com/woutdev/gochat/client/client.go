package main

import (
	"./clientcli"
	"./io"
	"log"
	"strconv"
)

func main() {
	address, port, username, err := clientcli.Init()

	if err {
		return
	}

	log.Println("Attempting to connect to", address+":"+strconv.Itoa(port), "as", username)

	io.Connect(address, port, username)
}
