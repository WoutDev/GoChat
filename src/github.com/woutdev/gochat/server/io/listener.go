package io

import (
	"fmt"
	"net"
	"strconv"
)

var clients []Client

func ListenForConnectionsForever(address string, port int) {
	portStr := strconv.Itoa(port)

	ln, _ := net.Listen("tcp", address+":"+portStr)

	fmt.Println("Now listening on", ln.Addr())

	for {
		conn, _ := ln.Accept()

		client := Client{conn}

		go client.Listen(&clients)

		clients = append(clients, client)
	}
}
