package io

import (
	"log"
	"net"
	"strconv"
)

var clients []Client

func ListenForConnectionsForever(address string, port int) {
	portStr := strconv.Itoa(port)

	ln, err := net.Listen("tcp", address+":"+portStr)

	if err != nil {
		log.Fatal("Error listening: ", err)
	}

	log.Println("Now listening!")

	for {
		conn, _ := ln.Accept()

		client := Client{conn}

		go client.Listen(&clients)

		clients = append(clients, client)
	}
}
