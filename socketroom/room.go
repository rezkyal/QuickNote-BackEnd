package socketroom

import (
	"github.com/gorilla/websocket"
	"github.com/rezkyal/QuickNote-BackEnd/queryfunction"
)

type Room struct {
	noteQuery *queryfunction.NoteQuery

	Forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.Forward:
			for client := range r.clients {
				client.send <- msg
			}
		}

	}
}

func (r *Room) AddClient(c *websocket.Conn) {
	client := &client{
		socket: c,
		send:   make(chan []byte, 1024),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func NewRoom(noteQuery *queryfunction.NoteQuery) *Room {
	return &Room{
		noteQuery: noteQuery,
		Forward:   make(chan []byte),
		join:      make(chan *client),
		leave:     make(chan *client),
		clients:   make(map[*client]bool),
	}
}
