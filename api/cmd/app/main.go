package main

import (
	"log"
	"net/http"

	"github.com/juliogsn/go-rt-auction/internal/handlers/auctions"
	"github.com/rs/cors"
)

func main() {
	hub := newHub()
	go hub.run()

  mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	mux.HandleFunc("/auction", auctions.List)

  handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
