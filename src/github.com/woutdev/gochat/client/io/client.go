package io

import (
	"../../protocol"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func Connect(address string, port int, username string) {
	portStr := strconv.Itoa(port)

	conn, _ := net.Dial("tcp", address+":"+portStr)

	connPacket := protocol.ConnectPacket{
		Id: protocol.CONNECT,
	}

	connPacket.SetUsername(username)

	connPacketEncoded, _ := protocol.Encode(&connPacket)

	fmt.Fprintf(conn, string(connPacketEncoded)+"\n")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected! Type anything and press return to send.")

	go listenForInput(&conn)

	go listenForOutput(&conn, reader)

	for {
		time.Sleep(time.Second)
	}
}

func listenForOutput(conn *net.Conn, reader *bufio.Reader) {
	for {
		text, _ := reader.ReadString('\n')

		p := protocol.MessagePacket{
			Id: protocol.MESSAGE,
		}

		p.SetMessage(text)

		encodedPacket, _ := protocol.Encode(&p)

		fmt.Fprintf(*conn, string(encodedPacket)+"\n")
	}
}
func listenForInput(conn *net.Conn) {
	for {
		message, err := bufio.NewReader(*conn).ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		p, err := protocol.Decode([]byte(message))

		fmt.Println("Server -> Us: ", p.GetId())
	}
}
