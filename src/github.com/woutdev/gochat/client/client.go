package main

import (
	"./clientcli"
	"./io"
	"fmt"
)

func main() {
	address, port, username, err := clientcli.Init()

	if err {
		return
	}

	fmt.Println("Attempting to connect to", address, ":", port, "as", username)

	io.Connect(address, port, username)
}
