package protocol

import (
	"encoding/json"
)

type MessagePacket struct {
	Id int8
	Msg string
}

func (c *MessagePacket) SetMessage(msg string) {
	c.Msg = msg
}

func (c *MessagePacket) GetMessage() string {
	return c.Msg
}

func (c *MessagePacket) GetId() int8 {
	return c.Id
}

func (c *MessagePacket) SetId(id int8) {
	c.Id = id
}

func (c *MessagePacket) Handle(data map[string] *json.RawMessage) {
	json.Unmarshal(*data["Msg"], &c.Msg)
}