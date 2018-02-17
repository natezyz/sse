package sse

import (
	"log"
)

type Channel struct {
	lastEvent string
	name      string
	clients   map[*Client]bool
}

func NewChannel(name string) (channel *Channel) {
	return &Channel{
		lastEvent: "",
		name:      name,
		clients:   make(map[*Client]bool),
	}
}

func (c *Channel) SendMessage(message *Message) {
	c.lastEvent = message.id

	for client, open := range c.clients {
		if open {
			client.send <- message
		}
	}

	log.Printf("Sent message to %d clients on '%s'.\n", len(c.clients), c.name)
}

func (c *Channel) Close() {
	for client, _ := range c.clients {
		c.remove(client)
	}

	log.Printf("closed channel '%s'\n", c.name)
}

func (c *Channel) Size() int {
	return len(c.clients)
}

func (c *Channel) LastEventId() string {
	return c.lastEvent
}

func (c *Channel) remove(client *Client) {
	close(client.send)
	delete(c.clients, client)
	log.Printf("client disconnected from channel '%s'\n", c.name)
}

func (c *Channel) addClient(client *Client) {
	c.clients[client] = true
	log.Printf("client added to channel '%s'\n", c.name)
}
