package io

import (
	"../../protocol"
	"bufio"
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func (client *Client) Listen(clients *[]Client) {
	for {
		msg, _ := bufio.NewReader(client.Conn).ReadString('\n')

		p, err := protocol.Decode([]byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		switch p.GetId() {
		case protocol.CONNECT:
			username := p.(*protocol.ConnectPacket).GetUsername()
			log.Println(client.Conn.RemoteAddr(), "CONNECT", " >>> ", "New connection with username: ", username)
		case protocol.MESSAGE:
			msg := p.(*protocol.MessagePacket).GetMessage()
			log.Println(client.Conn.RemoteAddr(), "MSG", " >>> ", "Incoming message: ", msg)
		default:
			log.Println(client.Conn.RemoteAddr(), p.GetId())
		}

		for _, c := range *clients {
			encodedPacket, _ := protocol.Encode(p)
			c.Conn.Write(encodedPacket)
		}
	}
}
