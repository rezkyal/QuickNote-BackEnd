package socketroom

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
}

type Message struct {
	Noteid    string    `json:"noteid"`
	Title     string    `json:"title"`
	Note      string    `json:"note"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()

		jsonmsg := Message{}

		// json.Unmarshal(msg, &jsonmsg)
		json.Unmarshal(msg, &jsonmsg)

		noteidint, err := strconv.ParseInt(jsonmsg.Noteid, 10, 64)
		if err != nil {
			log.Panic(err)
		}

		note := c.room.noteQuery.FindNote(noteidint)
		note.Title = jsonmsg.Title
		note.Note = jsonmsg.Note

		jsonmsg.Timestamp = time.Now().UTC()

		strjson, err := json.Marshal(jsonmsg)

		if err != nil {
			log.Panic(err)
		}

		c.room.noteQuery.UpdateNote(note)

		if err != nil {
			log.Panic(err)
		}

		c.room.Forward <- []byte(strjson)
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
