package main

import (
	"log"
	"net/http"

	"github.com/juliogsn/go-rt-auction/internal/handlers/auctions"
	"github.com/juliogsn/go-rt-auction/pkg/conn"
	"github.com/rs/cors"
)

func main() {
	hub := conn.NewHub()
	go hub.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client := conn.ServeWs(hub, w, r)
		log.Println("Client comes in")
		if client != nil {
			log.Println("Creating bidder")
			bidder := Bidder{
				Id:     1,
				Name:   "Júlio",
				Client: client,
			}
			bidder.connect()
			log.Println("bidder connected")
		}
	})

	mux.HandleFunc("/auction", auctions.List)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
