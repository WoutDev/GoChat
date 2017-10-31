package protocol

import (
	"encoding/json"
)

type MessagePacket struct {
	Id       int8
	Msg      string
	Username string
}

func (c *MessagePacket) SetMessage(msg string) {
	c.Msg = msg
}

func (c *MessagePacket) GetMessage() string {
	return c.Msg
}

func (c *MessagePacket) SetUsername(username string) {
	c.Username = username
}

func (c *MessagePacket) GetUsername() string {
	return c.Username
}

func (c *MessagePacket) GetId() int8 {
	return c.Id
}

func (c *MessagePacket) SetId(id int8) {
	c.Id = id
}

func (c *MessagePacket) Handle(data map[string]*json.RawMessage) {
	json.Unmarshal(*data["Msg"], &c.Msg)
	json.Unmarshal(*data["Username"], &c.Username)
}
