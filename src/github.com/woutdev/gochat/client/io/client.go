package io

import (
	"../../protocol"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var username string
var connected bool = true

func Connect(address string, port int, user string) {
	portStr := strconv.Itoa(port)

	conn, err := net.Dial("tcp", address+":"+portStr)

	if err != nil {
		log.Fatal("Error: failed to connect: ", err.Error())
	}

	connPacket := protocol.ConnectPacket{
		Id: protocol.CONNECT,
	}

	username = user

	connPacket.SetUsername(username)

	connPacketEncoded, _ := protocol.Encode(&connPacket)

	fmt.Fprintf(conn, string(connPacketEncoded)+"\n")

	reader := bufio.NewReader(os.Stdin)

	log.Println("Connected! Type anything and press return to send. Type /quit to quit")

	go listenForInput(&conn)

	go listenForOutput(&conn, reader)

	for {
		time.Sleep(time.Second)
	}
}

func listenForOutput(conn *net.Conn, reader *bufio.Reader) {
	for connected {
		text, _ := reader.ReadString('\n')

		if strings.TrimRight(text, "\r\n") == "/quit" {
			p := protocol.DisconnectPacket{
				Id: protocol.DISCONNECT,
			}

			p.SetUsername(username)

			encodedPacket, _ := protocol.Encode(&p)

			fmt.Fprint(*conn, string(encodedPacket)+"\n")

			connected = false

			os.Exit(1)

			break
		}

		p := protocol.MessagePacket{
			Id: protocol.MESSAGE,
		}

		p.SetMessage(text)
		p.SetUsername(username)

		encodedPacket, _ := protocol.Encode(&p)

		fmt.Fprintf(*conn, string(encodedPacket)+"\n")
	}
}
func listenForInput(conn *net.Conn) {
	for connected {
		message, err := bufio.NewReader(*conn).ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		p, err := protocol.Decode([]byte(message))

		fmt.Println("Server -> Us: ", p.GetId())
	}
}
