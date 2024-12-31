package main

import (
	"log"

	"github.com/juliogsn/go-rt-auction/pkg/conn"
)

type Bidder struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Client *conn.Client
}

func (b *Bidder) connect() {
	message := Message{
		Type: "bidder",
		Data: b,
	}

	err := b.Client.GetConn().WriteJSON(message)
	if err != nil {
		log.Println("error connecting bidder sending message")
		return
	}
}
