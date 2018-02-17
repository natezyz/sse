package sse

import (
	"log"
	"net/http"
)

type Server struct {
	channels     map[string]*Channel
	newClient    chan *Client
	removeClient chan *Client
	shutdown     chan bool
	closeChannel chan string
}

func NewServer() *Server {
	s := &Server{
		make(map[string]*Channel),
		make(chan *Client),
		make(chan *Client),
		make(chan bool),
		make(chan string),
	}

	log.Printf("starter sse server")
	go s.listen()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) SendMessage(channel string, message *Message) {
	if len(channel) == 0 {
		for name, channel := range s.channels {
			channel.SendMessage(message)
		}
	} else if channel, ok := s.channels[channel]; ok {
		channel.SendMessage(message)
	} else {
		log.Println("no message sent, channel '%s' is empty", channel)
	}
}

func (s *Server) Restart() {
	log.Printf("restarting server\n")
	s.close()
}

func (s *Server) Shutdown() {
	log.Printf("shutting down server\n")
	s.shutdown <- true
}

func (s *Server) Count() int {
	i := 0
	for _, ch := range s.channels() {
		i += ch.ClientCount()
	}
	return i
}

func (s *Server) Channels() []string {
	return make([]string)
}
