package protocol

import "encoding/json"

type ConnectPacket struct {
	Id int8
	Username string
}

func (c *ConnectPacket) SetUsername(username string) {
	c.Username = username
}

func (c *ConnectPacket) GetUsername() string {
	return c.Username
}

func (c *ConnectPacket) GetId() int8 {
	return c.Id
}

func (c *ConnectPacket) SetId(id int8) {
	c.Id = id
}

func (c *ConnectPacket) Handle(data map[string]*json.RawMessage) {
	json.Unmarshal(*data["Username"], &c.Username)
}