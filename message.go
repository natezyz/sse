package sse

import (
	"bytes"
	"log"
)

type Message struct {
	id,
	data,
	event string
	retry int
}

func NewMessage(id, data, event string) (msg *Message) {
	return &Message{
		id, data, event, 0,
	}
}

func (m *Message) String() string {
	var buffer bytes.Buffer

	if len(m.id) > 0 {
		buffer.WriteString(fmt.Sprintf("id: %s\n", m.id))
	}

	if len(m.data) > 0 {
		buffer.WriteString(fmt.Sprintf("data: %s\n", m.data))
	}

	if m.retry > 0 {
		buffer.WriteString(fmt.Sprintf("retry: %d\n", m.retry))
	}

	if len(m.event) > 0 {
		buffer.WriteString(fmt.Sprintf("event: %s\n", m.event))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
