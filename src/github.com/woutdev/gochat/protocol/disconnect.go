package protocol

import "encoding/json"

type DisconnectPacket struct {
	Id       int8
	Username string
}

func (c *DisconnectPacket) SetUsername(username string) {
	c.Username = username
}

func (c *DisconnectPacket) GetUsername() string {
	return c.Username
}

func (c *DisconnectPacket) GetId() int8 {
	return c.Id
}

func (c *DisconnectPacket) SetId(id int8) {
	c.Id = id
}

func (c *DisconnectPacket) Handle(data map[string]*json.RawMessage) {
	json.Unmarshal(*data["Username"], &c.Username)
}
