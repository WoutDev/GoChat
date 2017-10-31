package protocol

import (
	"encoding/json"
	"fmt"
)

const (
	CONNECT    = 0x01
	MESSAGE    = 0x02
	DISCONNECT = 0x03
)

type Packet interface {
	GetId() int8
	SetId(int8)
	Handle(map[string]*json.RawMessage)
}

func Decode(data []byte) (Packet, error) {
	var rawPacket map[string]*json.RawMessage

	err := json.Unmarshal(data, &rawPacket)

	if err != nil {
		e := fmt.Errorf("invalid packet, failed to decode packet")
		return nil, e
	}

	var pId int8

	err = json.Unmarshal(*rawPacket["Id"], &pId)

	if err != nil {
		e := fmt.Errorf("invalid packet id, failed to decode packet")
		return nil, e
	}

	var packet Packet

	switch pId {
	case CONNECT:
		packet = &ConnectPacket{}
	case MESSAGE:
		packet = &MessagePacket{}
	case DISCONNECT:
		packet = &DisconnectPacket{}
	default:
		packet = nil
	}

	if packet == nil {
		e := fmt.Errorf("invalid packet id, failed to find correct packet by id")
		return nil, e
	}

	packet.SetId(pId)

	packet.Handle(rawPacket)

	return packet, nil
}

func Encode(packet Packet) ([]byte, error) {
	return json.Marshal(packet)
}
