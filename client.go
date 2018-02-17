package sse

import (
	"log"
)

type Client struct {
	lastEvent string
	name      string
	send      chan *Message
}

func NewClient(lastEvent, name string) (client *Client) {
	return &Client{
		lastEvent: lastEvent,
		name:      name,
		send:      make(chan *Message),
	}
}

func (c *Client) SendMessage(message *Message) {
	c.lastEvent = message.id
	c.send <- message
}

func (c *Client) Channel() string {
	return c.name
}

func (c *Client) LastEvent() string {
	return c.lastEvent
}
